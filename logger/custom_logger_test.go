package logger

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"talaria/utils"
	"testing"
	"time"
)

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

func TestNewCustomeHandler(t *testing.T) {
	var buf bytes.Buffer
	type args struct {
		opts *Options
	}
	tests := []struct {
		name    string
		args    args
		want    *CustomHandler
		wantOut string
	}{
		{name: "Happy flow retriving INFO", want: NewCustomeHandler(&buf, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if got := NewCustomeHandler(out, tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomeHandler() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("NewCustomeHandler() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestCustomHandler_Handle(t *testing.T) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	var b bytes.Buffer

	recWithAttr := slog.NewRecord(time.Now(), slog.LevelInfo, "some message", pcs[0])
	recWithAttr.AddAttrs(slog.Attr{Key: "key1", Value: slog.AnyValue("value1")})

	type fields struct {
		opts           Options
		preformatted   []byte
		unopenedGroups []string
		indentLevel    int
		mu             *sync.Mutex
		out            io.Writer
	}
	type args struct {
		ctx context.Context
		r   slog.Record
	}
	tests := []struct {
		name                               string
		fields                             fields
		args                               args
		wantErr                            bool
		want, wantTime, wantLevel, wantMsg string
	}{
		{name: "Happy flow retriving INFO", fields: fields{mu: &sync.Mutex{}, out: &b}, args: args{r: slog.NewRecord(time.Now(), slog.LevelInfo, "some message", pcs[0])}, wantErr: false, want: "some message", wantTime: "time=", wantLevel: "level=", wantMsg: "msg="},
		{name: "Happy flow retriving INFO with attributes", fields: fields{mu: &sync.Mutex{}, out: &b}, args: args{r: recWithAttr}, wantErr: false, want: "value1", wantTime: "time=", wantLevel: "level=", wantMsg: "msg="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ch := &CustomHandler{
				opts:           tt.fields.opts,
				preformatted:   tt.fields.preformatted,
				unopenedGroups: tt.fields.unopenedGroups,
				indentLevel:    tt.fields.indentLevel,
				mu:             tt.fields.mu,
				out:            tt.fields.out,
			}
			if err := ch.Handle(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CustomHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := b.String()

			if !strings.Contains(got, tt.wantTime) {
				t.Errorf("Custome logger should contain %s", tt.wantTime)
			}

			if !strings.Contains(got, tt.wantLevel) {
				t.Errorf("Custome logger should contain %s", tt.wantLevel)
			}

			if !strings.Contains(got, tt.wantMsg) {
				t.Errorf("Custome logger should contain %s", tt.wantMsg)
			}

			if !strings.Contains(got, tt.want) {
				t.Errorf("Custome logger should contain %s", tt.want)
			}

		})
	}
}
