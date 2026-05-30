//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestChainContainsNewHono_None 테스트
package hono

import "testing"

func TestChainContainsNewHono_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = new Other();`+"\n"))
	if chainContainsNewHono(fi.Root, fi.Src) {
		t.Fatal("expected no new Hono()")
	}
}
