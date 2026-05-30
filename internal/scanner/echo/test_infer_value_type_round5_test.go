//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestInferValueType_Round5 테스트
package echo

import "testing"

func TestInferValueType_Round5(t *testing.T) {
	cases := map[string]string{
		`"s"`:   "string",
		`42`:    "integer",
		`3.14`:  "number",
		`true`:  "boolean",
		`false`: "boolean",
		`nil`:   "null",
	}
	for src, want := range cases {
		if got := inferValueType(parseExpr(t, src), nil); got != want {
			t.Errorf("inferValueType(%s)=%q want %q", src, got, want)
		}
	}
}
