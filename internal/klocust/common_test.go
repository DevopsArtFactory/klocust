package klocust

import "testing"

func Test_getLocustConfigFilename(t *testing.T) {
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
			want: "hello" + LocustConfigFileWithExtension,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLocustConfigFilename(tt.args.locustName); got != tt.want {
				t.Errorf("getLocustConfigFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			want: LocustMasterDeploymentPrefix + "hello",
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

func Test_getLocustFilename(t *testing.T) {
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
			want: "hello" + LocustFileWithExtension,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLocustFilename(tt.args.locustName); got != tt.want {
				t.Errorf("getLocustFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
