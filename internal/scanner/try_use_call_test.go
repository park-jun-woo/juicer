//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_ValidUse 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_ValidUse(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Use"}},
		Args: []ast.Expr{&ast.Ident{Name: "authMiddleware"}},
	}
	routers := map[string]*routerInfo{"r": {}}
	tryUseCall(call, routers)
	if len(routers["r"].middleware) != 1 {
		t.Fatal("expected 1 middleware")
	}
}

