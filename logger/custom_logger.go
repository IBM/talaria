package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
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
	h              slog.Handler
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
	fmt.Println("===>>> test111")
	indentLevel := 0
	bufp := allocBuf()
	buf := *bufp
	defer func() {
		*bufp = buf
		freeBuf(bufp)
	}()
	timestamp := time.Now().Format(time.RFC3339Nano)
	lev, colCode := colorLogLevel(r.Level.String())

	buf = append(buf, "time="...)
	buf = append(buf, painter(colCode, timestamp)...)
	buf = append(buf, " level="...)
	buf = append(buf, lev...)
	buf = append(buf, " msg="...)
	buf = append(buf, painter(colCode, r.Message)...)

	r.Attrs(func(a slog.Attr) bool {
		buf = ch.appendAttr(buf, a, indentLevel)
		return true
	})

	// logMessage := fmt.Sprintf("time=%s level=%s msg=%s", painter(green, timestamp), lev, painter(colCode, r.Message), painter(colCode, string(buf)))
	// fmt.Println(logMessage)

	// adding \n at the end for better formatting
	buf = append(buf, "\n"...)
	fmt.Println("===>>> test")
	ch.mu.Lock()
	defer ch.mu.Unlock()

	_, err := ch.out.Write(buf)
	if err != nil {
		fmt.Println("write out error ", err)
	}
	return err
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
	//return ch.h.Enabled(ctx, level)
	return level >= ch.opts.Level.Level()
}

func (ch *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// return &CustomHandler{h: ch.h.WithAttrs(attrs), b: ch.b, mu: ch.mu}
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
		chCopy.preformatted = chCopy.appendAttr(chCopy.preformatted, a, chCopy.indentLevel)
	}
	return &chCopy
}

func (ch *CustomHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
	for _, g := range ch.unopenedGroups {
		buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
		indentLevel++
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

func (ch *CustomHandler) appendAttr(buf []byte, a slog.Attr, colCode int) []byte {
	// Resolve the Attr's value before doing anything else
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs
	if a.Equal(slog.Attr{}) {
		return buf
	}

	buf = fmt.Appendf(buf, " %s=%s", a.Key, painter(colCode, a.Value.String()))
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
