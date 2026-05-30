//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_NotMethod 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_NotMethod(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{
		"app": {},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "app"},
			Sel: &ast.Ident{Name: "Use"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "logger"},
		},
	}

	_, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if ok {
		t.Fatal("expected false — Use is not a route method")
	}
}

func routeCallFrom(t *testing.T, src string) (*ast.CallExpr, *token.FileSet) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parserParseFile(fset, src)
	if err != nil {
		t.Fatal(err)
	}
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok && call == nil {
			if _, ok := c.Fun.(*ast.SelectorExpr); ok {
				call = c
			}
		}
		return true
	})
	return call, fset
}

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
	if len(exprs) != 2 { // mw + h
		t.Errorf("handlerExprs = %d, want 2", len(exprs))
	}
}

func TestTryRouteCall_NotSelector(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { plainCall() }\n")
	if call == nil {
		// no selector call found
		return
	}
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{}, "m.go", fset)
	if ok {
		t.Fatal("plain call should be false")
	}
}

func TestTryRouteCall_UnknownRecv(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { other.Get(\"/x\", h) }\n")
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{"app": {}}, "m.go", fset)
	if ok {
		t.Fatal("unknown receiver should be false")
	}
}

func TestTryRouteCall_TooFewArgs2(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { app.Get(\"/x\") }\n")
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{"app": {}}, "m.go", fset)
	if ok {
		t.Fatal("single arg should be false")
	}
}

func TestTryRouteCall_NonStringPath(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { app.Get(pathVar, h) }\n")
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{"app": {}}, "m.go", fset)
	if ok {
		t.Fatal("non-string path should be false")
	}
}
