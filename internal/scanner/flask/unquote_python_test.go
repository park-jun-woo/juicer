//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what unquotePython 테스트
package flask

import "testing"

func TestUnquotePython(t *testing.T) {
	tests := []struct{ in, want string }{
		{`"hello"`, "hello"},
		{`'hello'`, "hello"},
		{`f"/path"`, "/path"},
		{`r"\d+"`, `\d+`},
		{`f'/p'`, "/p"},
		{`r'\w'`, `\w`},
		{`"""triple"""`, "triple"},
		{`'''triple'''`, "triple"},
		{"x", "x"},
		{"", ""},
		{"ab", "ab"},
		{`"unterminated`, `"unterminated`},
	}
	for _, tt := range tests {
		if got := unquotePython(tt.in); got != tt.want {
			t.Errorf("unquotePython(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
