//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestIsNewHonoCall_NotNewExpr 테스트
package hono

import "testing"

func TestIsNewHonoCall_NotNewExpr(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	if isNewHonoCall(call, fi.Src) {
		t.Fatal("expected false for non-new_expression")
	}
}
