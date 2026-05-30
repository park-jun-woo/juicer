//ff:func feature=scan type=test control=sequence topic=hono
//ff:what chainContainsNewHono 테스트
package hono

import "testing"

func TestChainContainsNewHono(t *testing.T) {
	fi := mustParse(t, []byte(`const app = new Hono().get("/x", h);`+"\n"))
	if !chainContainsNewHono(fi.Root, fi.Src) {
		t.Fatal("expected new Hono() detected in chain")
	}
}

func TestChainContainsNewHono_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = new Other();`+"\n"))
	if chainContainsNewHono(fi.Root, fi.Src) {
		t.Fatal("expected no new Hono()")
	}
}
