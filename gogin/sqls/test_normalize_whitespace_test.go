//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what TestNormalizeWhitespace 테스트
package sqls

import (
	"testing"
)

func TestNormalizeWhitespace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"  hello   world  ", "hello world"},
		{"no change", "no change"},
		{"  \t  multiple  \n  spaces  ", "multiple spaces"},
	}
	for _, tt := range tests {
		got := normalizeWhitespace(tt.input)
		if got != tt.want {
			t.Errorf("normalizeWhitespace(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
