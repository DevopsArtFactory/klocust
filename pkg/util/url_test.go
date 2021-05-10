package util

import "testing"

func Test_IsURL(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"with correct URL",
			args{str: "https://www.example.com"},
			true,
		},
		{
			"with not URL",
			args{str: "it's not url"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsURL(tt.args.str); got != tt.want {
				t.Errorf("isURL(%s) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
