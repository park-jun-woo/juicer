//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestInferValueType_Literals 테스트
package fiber

import "testing"

func TestInferValueType_Literals(t *testing.T) {
	cases := map[string]string{
		`"x"`:      "string",
		"42":       "integer",
		"3.14":     "number",
		"true":     "boolean",
		"false":    "boolean",
		"nil":      "null",
		"[]int{1}": "array",
		"s[1:2]":   "array",
		"foo()":    "unknown",
	}
	for in, want := range cases {
		if got := inferFor(t, in); got != want {
			t.Errorf("inferValueType(%q) = %q, want %q", in, got, want)
		}
	}
}
