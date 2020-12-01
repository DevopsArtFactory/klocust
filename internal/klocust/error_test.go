/*
Copyright 2020 The klocust Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
			fields: fields{
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
			args: args{
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
