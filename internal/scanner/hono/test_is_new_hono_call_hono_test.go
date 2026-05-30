//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestIsNewHonoCall_Hono 테스트
package hono

import "testing"

func TestIsNewHonoCall_Hono(t *testing.T) {
	if _, ok := newExprOf(t, `const a = new Hono();`); !ok {
		t.Fatal("expected true for new Hono()")
	}
}
