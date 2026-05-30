//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestUnquoteTS 테스트
package hono

import "testing"

func TestUnquoteTS(t *testing.T) {
	cases := []struct{ in, want string }{
		{`"hi"`, "hi"},
		{`'hi'`, "hi"},
		{"`hi`", "hi"},
		{"x", "x"},
		{"", ""},
		{`"mismatch'`, `"mismatch'`},
		{"noquote", "noquote"},
	}
	for _, c := range cases {
		if got := unquoteTS(c.in); got != c.want {
			t.Errorf("unquoteTS(%q)=%q want %q", c.in, got, c.want)
		}
	}
}
