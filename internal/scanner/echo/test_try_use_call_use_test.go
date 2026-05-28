//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_Use 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_Use(t *testing.T) {
	routers := map[string]*routerInfo{
		"e": {},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "e"},
			Sel: &ast.Ident{Name: "Use"},
		},
		Args: []ast.Expr{&ast.Ident{Name: "authMiddleware"}},
	}
	tryUseCall(call, routers)
	if len(routers["e"].middleware) != 1 || routers["e"].middleware[0] != "authMiddleware" {
		t.Fatalf("expected [authMiddleware], got %v", routers["e"].middleware)
	}
}
