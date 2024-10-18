package utils

import "testing"

func TestTrimWhitespaces(t *testing.T) {
	type args struct {
		old string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Happy flow remove \n", args: args{old: "string with \n and nothing else"}, want: "stringwithandnothingelse"},
		{name: "Happy flow remove \t", args: args{old: "string with \t TAB and nothing else"}, want: "stringwithTABandnothingelse"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimWhitespaces(tt.args.old); got != tt.want {
				t.Errorf("TrimWhitespaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
