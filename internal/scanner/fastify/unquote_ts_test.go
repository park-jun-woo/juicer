//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what unquoteTS 테스트
package fastify

import "testing"

func TestUnquoteTS(t *testing.T) {
	tests := []struct{ in, want string }{
		{`"hello"`, "hello"},
		{`'hello'`, "hello"},
		{"`hello`", "hello"},
		{"x", "x"},          // too short
		{"", ""},            // empty
		{"ab", "ab"},        // len>=2 but unquoted
		{`"mismatch'`, `"mismatch'`}, // mismatched quotes
	}
	for _, tt := range tests {
		if got := unquoteTS(tt.in); got != tt.want {
			t.Errorf("unquoteTS(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
