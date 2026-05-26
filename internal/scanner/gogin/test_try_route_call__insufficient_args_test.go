//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_InsufficientArgs 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_InsufficientArgs(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/test"`}},
	}
	routers := map[string]*routerInfo{"r": {}}
	_, _, ok := tryRouteCall(call, routers, "", nil)
	if ok {
		t.Fatal("expected false with too few args")
	}
}
