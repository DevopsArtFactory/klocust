package klocust

import "testing"

func Test_getKLocustConfigFilename(t *testing.T) {
	type args struct {
		kLocustName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check return string",
			args: args{
				kLocustName: "hello",
			},
			want: "hello-klocust.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKLocustConfigFilename(tt.args.kLocustName); got != tt.want {
				t.Errorf("getKLocustConfigFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKLocustMasterDeploymentName(t *testing.T) {
	type args struct {
		kLocustName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check return string",
			args: args{
				kLocustName: "hello",
			},
			want: "locust-master-hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKLocustMasterDeploymentName(tt.args.kLocustName); got != tt.want {
				t.Errorf("getKLocustMasterDeploymentName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLocustFilename(t *testing.T) {
	type args struct {
		kLocustName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check return string",
			args: args{
				kLocustName: "hello",
			},
			want: "hello-locustfile.py",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLocustFilename(tt.args.kLocustName); got != tt.want {
				t.Errorf("getLocustFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
