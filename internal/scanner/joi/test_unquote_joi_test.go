//ff:func feature=scan type=test topic=joi control=iteration dimension=1
//ff:what unquoteJoi 따옴표(',",`) 제거 테스트
package joi

import "testing"

func TestUnquoteJoi(t *testing.T) {
	cases := []struct{ in, want string }{
		{`'a'`, "a"},
		{`"b"`, "b"},
		{"`c`", "c"},
		{"x", "x"},
		{"", ""},
		{`'mismatch"`, `'mismatch"`},
	}
	for _, c := range cases {
		if got := unquoteJoi(c.in); got != c.want {
			t.Errorf("unquoteJoi(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
