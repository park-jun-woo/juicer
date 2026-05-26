//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_ValidCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_ValidCov(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("test.go", -1, 100)
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/api/users/:id"`}, &ast.Ident{Name: "handler"}},
	}
	routers := map[string]*routerInfo{"r": {prefix: "", middleware: []string{"auth"}}}
	ep, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if !ok {
		t.Fatal("expected true")
	}
	if ep.Method != "GET" {
		t.Fatalf("expected GET, got %s", ep.Method)
	}
}
