//ff:func feature=scan type=test control=sequence
//ff:what TestRouterArgIndex_Found 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestRouterArgIndex_Found(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "RegisterHandlersWithOptions"},
		Args: []ast.Expr{
			&ast.Ident{Name: "router"},
			&ast.Ident{Name: "si"},
			&ast.Ident{Name: "opts"},
		},
	}
	idx := routerArgIndex(call, "router")
	if idx != 0 {
		t.Fatalf("expected 0, got %d", idx)
	}
}
