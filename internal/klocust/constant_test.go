package klocust

import "testing"

func Test_getLocustMasterDeploymentName(t *testing.T) {
	type args struct {
		locustName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check return string",
			args: args{
				locustName: "hello",
			},
			want: locustMasterDeploymentPrefix + "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLocustMasterDeploymentName(tt.args.locustName); got != tt.want {
				t.Errorf("getLocustMasterDeploymentName() = %v, want %v", got, tt.want)
			}
		})
	}
}
