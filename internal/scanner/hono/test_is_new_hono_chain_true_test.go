//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestIsNewHonoChain_True 테스트
package hono

import "testing"

func TestIsNewHonoChain_True(t *testing.T) {
	call, src := firstCallExpr(t, `new Hono().basePath("/api");`)
	if !isNewHonoChain(call, src) {
		t.Fatal("expected true for new Hono().basePath() chain")
	}
}
