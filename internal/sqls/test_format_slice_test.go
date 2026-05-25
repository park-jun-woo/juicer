//ff:func feature=ratchet type=session control=iteration dimension=1
//ff:what TestFormatSlice 테스트
package sqls

import (
	"testing"
)

func TestFormatSlice(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{"a", "b", "c"}, "[a, b, c]"},
		{[]string{}, "[]"},
		{[]string{"x"}, "[x]"},
	}

	for _, tt := range tests {
		got := formatSlice(tt.input)
		if got != tt.want {
			t.Errorf("formatSlice(%v) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
