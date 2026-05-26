//ff:func feature=scan type=test control=sequence
//ff:what TestRouterArgIndex_NonIdent 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestRouterArgIndex_NonIdent(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "fn"},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"hello"`}},
	}
	idx := routerArgIndex(call, "router")
	if idx != -1 {
		t.Fatalf("expected -1, got %d", idx)
	}
}
