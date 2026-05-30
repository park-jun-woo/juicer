//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what unquoteTS: 작은/큰/백틱 따옴표 제거 + 짧은문자열/비따옴표 분기
package express

import "testing"

func TestUnquoteTS(t *testing.T) {
	cases := []struct{ in, want string }{
		{"'abc'", "abc"},   // single quotes
		{`"abc"`, "abc"},   // double quotes
		{"`abc`", "abc"},   // backticks
		{"x", "x"},         // too short (<2)
		{"abc", "abc"},     // no surrounding quotes
		{"'mismatch\"", "'mismatch\""}, // mismatched quotes
		{"", ""},           // empty
	}
	for _, c := range cases {
		if got := unquoteTS(c.in); got != c.want {
			t.Errorf("unquoteTS(%q)=%q want %q", c.in, got, c.want)
		}
	}
}
