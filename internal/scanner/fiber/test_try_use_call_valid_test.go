//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_Valid 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_Valid(t *testing.T) {
	routers := map[string]*routerInfo{
		"app": {},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "app"},
			Sel: &ast.Ident{Name: "Use"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "logger"},
		},
	}

	tryUseCall(call, routers)
	if len(routers["app"].middleware) != 1 || routers["app"].middleware[0] != "logger" {
		t.Fatalf("expected [logger], got %v", routers["app"].middleware)
	}
}
