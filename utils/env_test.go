package utils

import (
	"os"
	"testing"
)

func TestGetEnvVar(t *testing.T) {
	type args struct {
		lookUpVar, defaultVal, envVarKeyToBeSet, envVarValToBeSet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Happy flow retriving env var", args: args{lookUpVar: "BROKER_HOST", defaultVal: "BROKER_HOST", envVarKeyToBeSet: "BROKER_HOST", envVarValToBeSet: "0.0.0.0"}, want: "0.0.0.0"},
		{name: "Retriving env var returns default value", args: args{lookUpVar: "DUMMY", defaultVal: "0.0.0.0", envVarKeyToBeSet: "BROKER_HOST", envVarValToBeSet: ""}, want: "0.0.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.args.envVarKeyToBeSet, tt.args.envVarValToBeSet)
			if got := GetEnvVar(tt.args.lookUpVar, tt.args.defaultVal); got != tt.want {
				t.Errorf("GetEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
