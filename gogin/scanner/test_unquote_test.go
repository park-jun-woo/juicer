//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestUnquote 테스트
package scanner

import (
	"testing"
)

func TestUnquote(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`"hello"`, "hello"},
		{`"/api/v1"`, "/api/v1"},
		{"`raw`", "raw"},
		{"noquotes", "noquotes"},
	}
	for _, tt := range tests {
		got := unquote(tt.input)
		if got != tt.want {
			t.Errorf("unquote(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
