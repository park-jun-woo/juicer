//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestTryRouteCall_Guards_Round5 테스트
package echo

import (
	"go/token"
	"testing"
)

func TestTryRouteCall_Guards_Round5(t *testing.T) {
	routers := map[string]*routerInfo{"e": {}}
	fset := token.NewFileSet()

	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `f("/x", h)`), routers, "m.go", fset); ok {
		t.Fatal("non-selector should fail")
	}

	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `e.NotAMethod("/x", h)`), routers, "m.go", fset); ok {
		t.Fatal("non-echo method should fail")
	}

	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `unknown.GET("/x", h)`), routers, "m.go", fset); ok {
		t.Fatal("unknown router should fail")
	}

	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `e.GET("/x")`), routers, "m.go", fset); ok {
		t.Fatal("too few args should fail")
	}

	ep, _, ok := tryRouteCall(nil, callExprFrom(t, `e.GET("/x", h)`), routers, "m.go", fset)
	if !ok || ep.Method != "GET" || ep.Path != "/x" {
		t.Fatalf("valid route: %+v %v", ep, ok)
	}
}
