//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestExprString 테스트
package fiber

import "testing"

func TestExprString(t *testing.T) {
	cases := map[string]string{
		"x":              "x",
		"pkg.Field":      "pkg.Field",
		"Book{}":         "Book{}",
		"*Book":          "*Book",
		"&req":           "req",
		"make()":         "make()",
		"m[k]":           "m[k]",
		"map[string]int": "map[string]int",
		"[]byte":         "[]byte",
		"interface{}":    "interface{}",
		"42":             "42",
	}
	for in, want := range cases {
		if got := exprStringFor(t, in); got != want {
			t.Errorf("exprString(%q) = %q, want %q", in, got, want)
		}
	}
}
