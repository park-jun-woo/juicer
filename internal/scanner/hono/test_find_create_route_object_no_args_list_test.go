//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindCreateRouteObject_NoArgsList 테스트
package hono

import "testing"

func TestFindCreateRouteObject_NoArgsList(t *testing.T) {

	call, _ := firstCallExpr(t, `app.openapi();`)
	if findCreateRouteObject(call) != nil {
		t.Fatal("expected nil")
	}
}
