//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_WithPrefix 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_WithPrefix(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{
		"api": {prefix: "/api/v1"},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "api"},
			Sel: &ast.Ident{Name: "Post"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"/users"`},
			&ast.Ident{Name: "createUser"},
		},
	}

	ep, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if !ok {
		t.Fatal("expected ok")
	}
	if ep.Method != "POST" {
		t.Fatalf("expected POST, got %s", ep.Method)
	}
	if ep.Path != "/api/v1/users" {
		t.Fatalf("expected /api/v1/users, got %s", ep.Path)
	}
}
