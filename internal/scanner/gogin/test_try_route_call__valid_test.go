//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_Valid 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_Valid(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/test"`}, &ast.Ident{Name: "handler"}},
	}
	routers := map[string]*routerInfo{"r": {}}
	fset := token.NewFileSet()
	ep, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if !ok {
		t.Fatal("expected true")
	}
	if ep.Method != "GET" || ep.Path != "/test" {
		t.Fatalf("unexpected ep: %+v", ep)
	}
}
