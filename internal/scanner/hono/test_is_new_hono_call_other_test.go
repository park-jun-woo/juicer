//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestIsNewHonoCall_Other 테스트
package hono

import "testing"

func TestIsNewHonoCall_Other(t *testing.T) {
	if _, ok := newExprOf(t, `const a = new Other();`); ok {
		t.Fatal("expected false for new Other()")
	}
}
