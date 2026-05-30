//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveStatusCode_UnknownNilInfo 테스트
package echo

import "testing"

func TestResolveStatusCode_UnknownNilInfo(t *testing.T) {
	if got := resolveStatusCode(parseExpr(t, "someVar"), nil); got != "(unknown)" {
		t.Fatalf("got %q", got)
	}
}
