//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_Valid 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_Valid(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{
		"app": {prefix: ""},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "app"},
			Sel: &ast.Ident{Name: "Get"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"/users/:id"`},
			&ast.Ident{Name: "getUser"},
		},
	}

	ep, exprs, ok := tryRouteCall(call, routers, "main.go", fset)
	if !ok {
		t.Fatal("expected ok")
	}
	if ep.Method != "GET" {
		t.Fatalf("expected GET, got %s", ep.Method)
	}
	if ep.Path != "/users/:id" {
		t.Fatalf("expected /users/:id, got %s", ep.Path)
	}
	if len(exprs) != 1 {
		t.Fatalf("expected 1 handler expr, got %d", len(exprs))
	}
}
