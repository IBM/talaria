package main

import (
	"opentalaria/utils"
	"reflect"
	"testing"
)

func Test_parseListener(t *testing.T) {
	type args struct {
		l           string
		securityMap string
	}
	tests := []struct {
		name    string
		args    args
		want    Listener
		wantErr bool
	}{
		{
			name: "listener with ssl schema and empty host",
			args: args{
				l:           "SSL://:9092",
				securityMap: "",
			},
			want: Listener{
				Host:             "",
				Port:             9092,
				SecurityProtocol: utils.SSL,
				ListenerName:     "ssl",
			},
			wantErr: false,
		},
		{
			name: "listener with plaintext schema and localhost",
			args: args{
				l:           "PLAINTEXT://localhost:9092",
				securityMap: "",
			},
			want: Listener{
				Host:             "localhost",
				Port:             9092,
				SecurityProtocol: utils.PLAINTEXT,
				ListenerName:     "plaintext",
			},
			wantErr: false,
		},
		{
			name: "custom listener name",
			args: args{
				l:           "CUSTOM://localhost:9092",
				securityMap: "CUSTOM:PLAINTEXT",
			},
			want: Listener{
				Host:             "localhost",
				Port:             9092,
				SecurityProtocol: utils.PLAINTEXT,
				ListenerName:     "custom",
			},
			wantErr: false,
		},
		{
			name: "custom listener name not in security map",
			args: args{
				l:           "CUSTOM://localhost:9092",
				securityMap: "",
			},
			want:    Listener{},
			wantErr: true,
		},
		{
			name: "incorrect security protocol in security map",
			args: args{
				l:           "CUSTOM://localhost:9092",
				securityMap: "CUSTOM:CUSTOM",
			},
			want:    Listener{},
			wantErr: true,
		},
		{
			name: "empty port",
			args: args{
				l:           "CUSTOM://localhost",
				securityMap: "",
			},
			want:    Listener{},
			wantErr: true,
		},
		{
			name: "invalid port",
			args: args{
				l:           "CUSTOM://localhost:aaaa",
				securityMap: "",
			},
			want:    Listener{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("listener.security.protocol.map", tt.args.securityMap)

			got, err := parseListener(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
