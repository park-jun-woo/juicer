//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_NonStringPath 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_NonStringPath(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.Ident{Name: "pathVar"}, &ast.Ident{Name: "handler"}},
	}
	routers := map[string]*routerInfo{"r": {}}
	fset := token.NewFileSet()
	_, _, ok := tryRouteCall(call, routers, "test.go", fset)
	if ok {
		t.Error("expected not ok for non-string path")
	}
}
