//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestStringLitValue_Round5 테스트
package echo

import "testing"

func TestStringLitValue_Round5(t *testing.T) {
	if got := stringLitValue(parseExpr(t, `"abc"`)); got != "abc" {
		t.Fatalf("got %q", got)
	}
	if got := stringLitValue(parseExpr(t, `42`)); got != "" {
		t.Fatalf("non-string: %q", got)
	}
	if got := stringLitValue(parseExpr(t, `foo`)); got != "" {
		t.Fatalf("ident: %q", got)
	}
}
