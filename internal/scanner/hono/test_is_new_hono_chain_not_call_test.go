//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestIsNewHonoChain_NotCall 테스트
package hono

import "testing"

func TestIsNewHonoChain_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`+"\n"))
	id := findAllByType(fi.Root, "identifier")[0]
	if isNewHonoChain(id, fi.Src) {
		t.Fatal("expected false for non-call node")
	}
}
