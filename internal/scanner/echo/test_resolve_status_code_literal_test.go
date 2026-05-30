//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveStatusCode_Literal 테스트
package echo

import "testing"

func TestResolveStatusCode_Literal(t *testing.T) {
	if got := resolveStatusCode(parseExpr(t, "200"), nil); got != "200" {
		t.Fatalf("got %q", got)
	}
}
