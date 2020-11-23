package util

import (
	"testing"
)

func TestIsFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "with exist file",
			args: args{filename: "fileutil_test.go"},
			want: true,
		},
		{
			name: "with not exist file",
			args: args{filename: "not_exist_file.txt"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFileExists(tt.args.filename); got != tt.want {
				t.Errorf("IsFileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
