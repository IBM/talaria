package main

import (
	"opentalaria/utils"
	"reflect"
	"testing"
)

func Test_parseBrokerName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   utils.SecurityProtocol
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseBrokerName(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBrokerName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseBrokerName() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseBrokerName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
