package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"slices"
	"strconv"
	"sync"
	"time"
)

const (
	noColor = "\033[0m"

	info  = "INFO"
	debug = "DEBUG"
	err   = "ERROR"
	warn  = "WARN"

	green  = 32
	yellow = 33
	cyan   = 36
	gray   = 37
	red    = 91
	white  = 97
)

type CustomHandler struct {
	opts           Options
	preformatted   []byte   // data from WithGroup and WithAttrs
	unopenedGroups []string // groups from WithGroup that haven't been opened
	indentLevel    int
	mu             *sync.Mutex
	out            io.Writer
}

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

func NewCustomeHandler(out io.Writer, opts *Options) *CustomHandler {
	h := &CustomHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

// CustomHandler is a custom handler for logging records.
func (ch *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	indentLevel := 0
	bufp := allocBuf()
	buf := *bufp
	defer func() {
		*bufp = buf
		freeBuf(bufp)
	}()
	lev, colCode := colorLogLevel(r.Level.String())

	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		// Optimize to minimize allocation.
		srcbufp := allocBuf()
		defer freeBuf(srcbufp)
		*srcbufp = append(*srcbufp, f.File...)
		*srcbufp = append(*srcbufp, ':')
		*srcbufp = strconv.AppendInt(*srcbufp, int64(f.Line), 10)
		buf = ch.appendAttr(buf, slog.String(slog.SourceKey, string(*srcbufp)), colCode, 0)
	}

	buf = formatLoggerOutput(buf, lev, r.Message, colCode)

	if r.NumAttrs() > 0 {
		buf = ch.appendUnopenedGroups(buf, ch.indentLevel)
		r.Attrs(func(a slog.Attr) bool {
			buf = ch.appendAttr(buf, a, colCode, indentLevel)
			return true
		})
	}

	// adding \n at the end for better formatting
	buf = append(buf, "\n"...)
	ch.mu.Lock()
	defer ch.mu.Unlock()
	_, err := ch.out.Write(buf)
	if err != nil {
		fmt.Println("write out error ", err)
	}
	return err
}

func formatLoggerOutput(buf []byte, lev, msg string, colCode int) []byte {
	timestamp := time.Now().Format(time.RFC3339Nano)
	buf = append(buf, "time="...)
	buf = append(buf, painter(colCode, timestamp)...)
	buf = append(buf, " level="...)
	buf = append(buf, lev...)
	buf = append(buf, " msg="...)
	buf = append(buf, painter(colCode, msg)...)
	return buf
}

// Painter is a function that takes in a Bash color code and a string, and returns a string with the given string painted in the specified color.
func painter(colorCode int, msg string) string {
	//formatting message with ANSI escape sequence and selected color
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), msg, noColor)
}

func colorLogLevel(level string) (string, int) {
	if level == info {
		return painter(green, info), green
	} else if level == debug {
		return painter(white, debug), gray
	} else if level == err {
		return painter(red, err), red
	} else if level == warn {
		return painter(yellow, warn), yellow
	} else {
		return painter(white, info), gray
	}

}

func (ch *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= ch.opts.Level.Level()
}

func (ch *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return ch
	}
	chCopy := *ch
	// Force an append to copy the underlying array.
	pre := slices.Clip(ch.preformatted)
	// Add all groups from WithGroup that haven't already been added.
	chCopy.preformatted = chCopy.appendUnopenedGroups(pre, chCopy.indentLevel)
	// Each of those groups increased the indent level by 1.
	chCopy.indentLevel += len(chCopy.unopenedGroups)
	// Now all groups have been opened.
	chCopy.unopenedGroups = nil
	// Pre-format the attributes.
	for _, a := range attrs {
		//TODO replace gray with required color val
		chCopy.preformatted = chCopy.appendAttr(chCopy.preformatted, a, gray, chCopy.indentLevel)
	}
	return &chCopy
}

func (ch *CustomHandler) appendAttr(buf []byte, a slog.Attr, colCode, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs
	if a.Equal(slog.Attr{}) {
		return buf
	}

	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = strconv.AppendQuote(buf, a.Value.String())
		//buf = append(buf, ' ')
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = a.Value.Time().AppendFormat(buf, time.RFC3339Nano)
		//buf = append(buf, ' ')
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = ch.appendAttr(buf, ga, colCode, indentLevel)
		}
		//buf = append(buf, ' ')
	default:
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = append(buf, a.Value.String()...)
		buf = append(buf, '\n')
	}
	return buf
}

func (ch *CustomHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return ch
	}
	chCopy := *ch
	// Add an unopened group to chCopy without modifying h.
	chCopy.unopenedGroups = make([]string, len(ch.unopenedGroups)+1)
	copy(chCopy.unopenedGroups, ch.unopenedGroups)
	chCopy.unopenedGroups[len(chCopy.unopenedGroups)-1] = name
	return &chCopy
}

func (ch *CustomHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
	for _, g := range ch.unopenedGroups {
		buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
		indentLevel++
	}
	return buf
}

func allocBuf() *[]byte {
	return bufPool.Get().(*[]byte)
}

var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 1024)
		return &b
	},
}

func freeBuf(b *[]byte) {
	// To reduce peak allocation, return only smaller buffers to the pool.
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}
