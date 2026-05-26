//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_TooFewArgsCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_TooFewArgsCov(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/api"`}},
	}
	routers := map[string]*routerInfo{"r": {}}
	_, _, ok := tryRouteCall(call, routers, "main.go", nil)
	if ok {
		t.Fatal("expected false")
	}
}
