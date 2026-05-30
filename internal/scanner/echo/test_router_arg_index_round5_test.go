//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestRouterArgIndex_Round5 테스트
package echo

import "testing"

func TestRouterArgIndex_Round5(t *testing.T) {
	call := callExprFrom(t, `register(e, g, x)`)
	if got := routerArgIndex(call, "g"); got != 1 {
		t.Fatalf("got %d", got)
	}
	if got := routerArgIndex(call, "missing"); got != -1 {
		t.Fatalf("missing: %d", got)
	}
}
