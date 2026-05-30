//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what unquoteRust — 문자열 리터럴 따옴표 제거 분기를 검증
package actix

import "testing"

func TestUnquoteRust(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{`"hello"`, "hello"}, // quoted
		{`hello`, "hello"},   // not quoted, len >= 2
		{`x`, "x"},           // len < 2
		{``, ""},             // empty, len < 2
		{`"only-open`, `"only-open`}, // missing closing quote -> unchanged
	}
	for _, c := range cases {
		if got := unquoteRust(c.in); got != c.want {
			t.Errorf("unquoteRust(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
