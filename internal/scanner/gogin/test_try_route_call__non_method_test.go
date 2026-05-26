//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_NonMethod 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryRouteCall_NonMethod(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Render"}},
	}
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{}, "", nil)
	if ok {
		t.Fatal("expected false")
	}
}
