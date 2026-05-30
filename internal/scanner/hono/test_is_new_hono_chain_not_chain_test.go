//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestIsNewHonoChain_NotChain 테스트
package hono

import "testing"

func TestIsNewHonoChain_NotChain(t *testing.T) {
	call, src := firstCallExpr(t, `app.get("/x", h);`)
	if isNewHonoChain(call, src) {
		t.Fatal("expected false for plain method call")
	}
}
