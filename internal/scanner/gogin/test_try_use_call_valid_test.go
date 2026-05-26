//ff:func feature=scan type=extract control=sequence
//ff:what TestTryUseCall_Valid 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_Valid(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Use"}},
		Args: []ast.Expr{&ast.Ident{Name: "auth"}},
	}
	routers := map[string]*routerInfo{"r": {}}
	tryUseCall(call, routers)
	if len(routers["r"].middleware) != 1 {
		t.Errorf("expected 1 middleware, got %d", len(routers["r"].middleware))
	}
}
