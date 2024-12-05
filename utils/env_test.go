package utils

import (
	"os"
	"testing"
)

func TestGetEnvVar(t *testing.T) {
	type args struct {
		lookUpVar, defaultVal, envVarKeyToBeSet, envVarValToBeSet string
	}

	type Want struct {
		envVar string
		wasSet bool
	}
	tests := []struct {
		name string
		args args
		want Want
	}{
		{name: "Happy flow retriving env var", args: args{lookUpVar: "BROKER_HOST", defaultVal: "BROKER_HOST", envVarKeyToBeSet: "BROKER_HOST", envVarValToBeSet: "0.0.0.0"}, want: Want{envVar: "0.0.0.0", wasSet: true}},
		{name: "Retriving env var returns default value", args: args{lookUpVar: "DUMMY", defaultVal: "0.0.0.0", envVarKeyToBeSet: "BROKER_HOST", envVarValToBeSet: ""}, want: Want{envVar: "0.0.0.0", wasSet: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.args.envVarKeyToBeSet, tt.args.envVarValToBeSet)
			if got, wasSet := GetEnvVar(tt.args.lookUpVar, tt.args.defaultVal); got != tt.want.envVar || wasSet != tt.want.wasSet {
				t.Errorf("GetEnvVar() = %v, want %v", got, tt.want.envVar)
				t.Errorf("GetEnvVar() = %v, want %v", wasSet, tt.want.wasSet)
			}
		})
	}
}
