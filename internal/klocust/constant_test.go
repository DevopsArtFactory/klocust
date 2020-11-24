package klocust

import "testing"

func Test_getLocustMainDeploymentName(t *testing.T) {
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
			want: locustMainDeploymentPrefix + "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLocustMainDeploymentName(tt.args.locustName); got != tt.want {
				t.Errorf("getLocustMainDeploymentName() = %v, want %v", got, tt.want)
			}
		})
	}
}
