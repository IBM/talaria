package logger

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
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
	h slog.Handler
	b *bytes.Buffer
	m *sync.Mutex
}

// CustomHandler is a custom handler for logging records.
func (h *CustomHandler) Handle(c context.Context, r slog.Record) error {
	timestamp := time.Now().Format(time.RFC3339Nano)
	lev, colCode := colorLogLevel(r.Level.String())
	logMessage := fmt.Sprintf("time=%s level=%s msg=%s", painter(green, timestamp), lev, painter(colCode, r.Message))

	fmt.Println(logMessage)
	return nil
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
	return ch.h.Enabled(ctx, level)
}

func (ch *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CustomHandler{h: ch.h.WithAttrs(attrs), b: ch.b, m: ch.m}
}

func (ch *CustomHandler) WithGroup(name string) slog.Handler {
	return &CustomHandler{h: ch.h.WithGroup(name), b: ch.b, m: ch.m}
}
