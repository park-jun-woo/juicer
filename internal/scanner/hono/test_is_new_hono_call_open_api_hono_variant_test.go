//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestIsNewHonoCall_OpenAPIHonoVariant 테스트
package hono

import "testing"

func TestIsNewHonoCall_OpenAPIHonoVariant(t *testing.T) {
	if _, ok := newExprOf(t, `const a = new OpenAPIHono();`); !ok {
		t.Fatal("expected true for new OpenAPIHono()")
	}
}
