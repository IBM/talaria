package logger

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"opentalaria/utils"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestCustomHandler_Handle(t *testing.T) {
	type fields struct {
		h slog.Handler
		b *bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		c context.Context
		r slog.Record
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTime  string
		wantLevel string
		wantMsg   string
	}{
		{name: "Happy flow retriving INFO", wantTime: "time=", wantLevel: "level=", wantMsg: "msg="},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			h := &CustomHandler{
				h: tt.fields.h,
				b: tt.fields.b,
				m: tt.fields.m,
			}

			copyStdout := os.Stdout // keeping copy of the real stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			h.Handle(tt.args.c, tt.args.r)
			w.Close()
			// test coverage errors if we don't restore stdout
			os.Stdout = copyStdout // restoring copy stdout
			got, _ := io.ReadAll(r)

			gotString := string(got[:])

			if !strings.Contains(gotString, tt.wantTime) {
				t.Errorf("Custome logger should contain %s", tt.wantTime)
			}

			if !strings.Contains(gotString, tt.wantLevel) {
				t.Errorf("Custome logger should contain %s", tt.wantLevel)
			}

			if !strings.Contains(gotString, tt.wantMsg) {
				t.Errorf("Custome logger should contain %s", tt.wantMsg)
			}

		})
	}
}

func Test_painter(t *testing.T) {
	type args struct {
		colorCode int
		msg       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Happy flow retriving INFO", args: args{colorCode: 32, msg: "dummy message"}, want: "[32mdummymessage[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := painter(tt.args.colorCode, tt.args.msg)
			got = utils.TrimWhitespaces(got)
			if got != tt.want {
				t.Errorf("painter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_colorLogLevel(t *testing.T) {
	const (
		info  = "INFO"
		debug = "DEBUG"
		err   = "ERROR"
		warn  = "WARN"
	)
	type args struct {
		level string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{name: "Happy flow retriving INFO", args: args{level: info}, want: "[32mINFO[0m", want1: 32},
		{name: "Happy flow retriving DEBUG", args: args{level: debug}, want: "[97mDEBUG[0m", want1: 37},
		{name: "Happy flow retriving INFO", args: args{level: err}, want: "[91mERROR[0m", want1: 91},
		{name: "Happy flow retriving INFO", args: args{level: warn}, want: "[33mWARN[0m", want1: 33},
		{name: "Happy flow retriving default case", args: args{level: "default"}, want: "[97mINFO[0m", want1: 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := colorLogLevel(tt.args.level)
			if got != tt.want {
				t.Errorf("colorLogLevel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("colorLogLevel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
