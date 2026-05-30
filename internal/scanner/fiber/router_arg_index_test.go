//ff:func feature=scan type=test control=sequence
//ff:what routerArgIndex — 라우터 인자 인덱스 테스트
package fiber

import "testing"

func TestRouterArgIndex(t *testing.T) {
	call := parseCall(t, "register(opts, app, more)")
	if got := routerArgIndex(call, "app"); got != 1 {
		t.Fatalf("app index = %d, want 1", got)
	}
	if got := routerArgIndex(call, "missing"); got != -1 {
		t.Fatalf("missing index = %d, want -1", got)
	}
}
