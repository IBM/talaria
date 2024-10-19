package logger

import (
	"bytes"
	"fmt"
	"reflect"
	"talaria/utils"
	"testing"
	"testing/slogtest"

	"gopkg.in/yaml.v2"
)

func TestSlogtest(t *testing.T) {
	var buf bytes.Buffer
	err := slogtest.TestHandler(NewCustomeHandler(&buf, nil), func() []map[string]any {
		return parseLogEntries(t, buf.Bytes())
	})
	if err != nil {
		t.Error(err)
	}
}

func parseLogEntries(t *testing.T, data []byte) []map[string]any {
	entries := bytes.Split(data, []byte("---\n"))
	entries = entries[:len(entries)-1] // last one is empty
	var ms []map[string]any
	for _, e := range entries {
		var m map[string]any
		if err := yaml.Unmarshal([]byte(e), &m); err != nil {
			t.Fatal(err)
		}
		ms = append(ms, m)
	}
	return ms
}

func TestParseLogEntries(t *testing.T) {
	in := `
a: 1
b: 2
c: 3
g:
    h: 4
    i: five
d: 6
---
e: 7
---
`
	want := []map[string]any{
		{
			"a": 1,
			"b": 2,
			"c": 3,
			"g": map[string]any{
				"h": 4,
				"i": "five",
			},
			"d": 6,
		},
		{
			"e": 7,
		},
	}
	got := parseLogEntries(t, []byte(in[1:]))

	fmt.Printf("%#v\n", got)
	fmt.Printf("%#v\n", want)
	fmt.Println(reflect.DeepEqual(got, want))

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
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
