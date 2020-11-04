package klocust

import (
	"reflect"
	"testing"
)

func TestFileExistsError_Error(t *testing.T) {
	type fields struct {
		Filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "check error messages",
			fields: fields {
				Filename: "test_file.txt",
			},
			want: "`test_file.txt` file is already exists.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := FileExistsError{
				Filename: tt.fields.Filename,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFileExistsError(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want FileExistsError
	}{
		{
			name: "create new error",
			args: args {
				filename: "test_file.txt",
			},
			want: FileExistsError{
				Filename: "test_file.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileExistsError(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileExistsError() = %v, want %v", got, tt.want)
			}
		})
	}
}
