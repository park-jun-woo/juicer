//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindCreateRouteObject_FirstNotCall 테스트
package hono

import "testing"

func TestFindCreateRouteObject_FirstNotCall(t *testing.T) {
	call, _ := firstCallExpr(t, `app.openapi(routeDef, h);`)
	if findCreateRouteObject(call) != nil {
		t.Fatal("expected nil for non-call first arg")
	}
}
