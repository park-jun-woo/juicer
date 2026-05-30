//ff:func feature=scan type=test control=sequence topic=hono
//ff:what isNewHonoChain 테스트
package hono

import "testing"

func TestIsNewHonoChain_True(t *testing.T) {
	call, src := firstCallExpr(t, `new Hono().basePath("/api");`)
	if !isNewHonoChain(call, src) {
		t.Fatal("expected true for new Hono().basePath() chain")
	}
}

func TestIsNewHonoChain_NotChain(t *testing.T) {
	call, src := firstCallExpr(t, `app.get("/x", h);`)
	if isNewHonoChain(call, src) {
		t.Fatal("expected false for plain method call")
	}
}

func TestIsNewHonoChain_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;` + "\n"))
	id := findAllByType(fi.Root, "identifier")[0]
	if isNewHonoChain(id, fi.Src) {
		t.Fatal("expected false for non-call node")
	}
}
