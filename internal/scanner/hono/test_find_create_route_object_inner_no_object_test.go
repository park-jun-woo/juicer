//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindCreateRouteObject_InnerNoObject 테스트
package hono

import "testing"

func TestFindCreateRouteObject_InnerNoObject(t *testing.T) {

	call, _ := firstCallExpr(t, `app.openapi(createRoute(x), h);`)
	if findCreateRouteObject(call) != nil {
		t.Fatal("expected nil when no object literal inside")
	}
}
