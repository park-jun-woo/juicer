//ff:func feature=scan type=test control=sequence
//ff:what TestRouterArgIndex_NotFound 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestRouterArgIndex_NotFound(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "doSomething"},
		Args: []ast.Expr{&ast.Ident{Name: "x"}},
	}
	idx := routerArgIndex(call, "router")
	if idx != -1 {
		t.Fatalf("expected -1, got %d", idx)
	}
}
