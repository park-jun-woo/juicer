//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_Success 테스트
package fiber

import "testing"

func TestTryRouteCall_Success(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { app.Get(\"/users/:id\", mw, h) }\n")
	routers := map[string]*routerInfo{"app": {prefix: "/api", middleware: []string{"auth"}}}
	ep, exprs, ok := tryRouteCall(call, routers, "main.go", fset)
	if !ok {
		t.Fatal("expected route match")
	}
	if ep.Method != "GET" || ep.Path != "/api/users/:id" {
		t.Fatalf("ep = %+v", ep)
	}
	if len(ep.Middleware) != 1 {
		t.Errorf("middleware not propagated: %v", ep.Middleware)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Errorf("path params not set: %+v", ep.Request)
	}
	if len(exprs) != 2 {
		t.Errorf("handlerExprs = %d, want 2", len(exprs))
	}
}
