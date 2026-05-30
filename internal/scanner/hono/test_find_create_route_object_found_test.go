//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindCreateRouteObject_Found 테스트
package hono

import "testing"

func TestFindCreateRouteObject_Found(t *testing.T) {
	call, _ := firstCallExpr(t, `app.openapi(createRoute({ method: "get" }), h);`)
	if findCreateRouteObject(call) == nil {
		t.Fatal("expected object node")
	}
}
