//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_NotGinMethod 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_NotGinMethod(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Use"}},
	}
	routers := map[string]*routerInfo{"r": {}}
	fset := token.NewFileSet()
	_, _, ok := tryRouteCall(call, routers, "test.go", fset)
	if ok {
		t.Error("expected not ok for Use")
	}
}
