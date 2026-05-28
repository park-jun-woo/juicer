//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_NotUse 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NotUse(t *testing.T) {
	routers := map[string]*routerInfo{
		"app": {},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "app"},
			Sel: &ast.Ident{Name: "Get"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "handler"},
		},
	}

	tryUseCall(call, routers)
	if len(routers["app"].middleware) != 0 {
		t.Fatalf("expected empty, got %v", routers["app"].middleware)
	}
}
