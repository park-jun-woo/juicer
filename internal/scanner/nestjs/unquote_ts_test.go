//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what unquoteTS 테스트
package nestjs

import "testing"

func TestUnquoteTS_Quotes(t *testing.T) {
	cases := []struct{ in, want string }{
		{`"hello"`, "hello"},
		{`'hello'`, "hello"},
		{"`hello`", "hello"},
		{"x", "x"},
		{"", ""},
		{`"a`, `"a`},
	}
	for _, c := range cases {
		got := unquoteTS(c.in)
		if got != c.want {
			t.Errorf("unquoteTS(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
