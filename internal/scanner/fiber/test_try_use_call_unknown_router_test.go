//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_UnknownRouter 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_UnknownRouter(t *testing.T) {
	routers := map[string]*routerInfo{}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "unknown"},
			Sel: &ast.Ident{Name: "Use"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "logger"},
		},
	}

	// Should not panic
	tryUseCall(call, routers)
}
