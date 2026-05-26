//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_NoRouterCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_NoRouterCov(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/api"`}, &ast.Ident{Name: "handler"}},
	}
	_, ok := tryRouteCall(call, map[string]*routerInfo{}, "", nil)
	if ok {
		t.Fatal("expected false")
	}
}
