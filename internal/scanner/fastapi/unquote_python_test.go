//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what unquotePython 테스트
package fastapi

import "testing"

func TestUnquotePython(t *testing.T) {
	tests := []struct{ in, want string }{
		{`"hello"`, "hello"},
		{`'hello'`, "hello"},
		{`f"/path"`, "/path"},
		{`r"\d+"`, `\d+`},
		{`"""triple"""`, "triple"},
		{`'''triple'''`, "triple"},
		{"x", "x"},
		{"", ""},
	}
	for _, tt := range tests {
		got := unquotePython(tt.in)
		if got != tt.want {
			t.Errorf("unquotePython(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
