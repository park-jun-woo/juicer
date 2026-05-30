//ff:func feature=scan type=test control=sequence topic=hono
//ff:what findCreateRouteObject 테스트
package hono

import "testing"

func TestFindCreateRouteObject_Found(t *testing.T) {
	call, _ := firstCallExpr(t, `app.openapi(createRoute({ method: "get" }), h);`)
	if findCreateRouteObject(call) == nil {
		t.Fatal("expected object node")
	}
}

func TestFindCreateRouteObject_NoArgsList(t *testing.T) {
	// no arguments at all -> handled via len check; use zero-arg call
	call, _ := firstCallExpr(t, `app.openapi();`)
	if findCreateRouteObject(call) != nil {
		t.Fatal("expected nil")
	}
}

func TestFindCreateRouteObject_FirstNotCall(t *testing.T) {
	call, _ := firstCallExpr(t, `app.openapi(routeDef, h);`)
	if findCreateRouteObject(call) != nil {
		t.Fatal("expected nil for non-call first arg")
	}
}

func TestFindCreateRouteObject_InnerNoObject(t *testing.T) {
	// createRoute called with no object literal -> nil
	call, _ := firstCallExpr(t, `app.openapi(createRoute(x), h);`)
	if findCreateRouteObject(call) != nil {
		t.Fatal("expected nil when no object literal inside")
	}
}
