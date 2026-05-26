//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_UnaryExpr 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExprString_UnaryExpr(t *testing.T) {
	got := exprString(&ast.UnaryExpr{Op: token.AND, X: &ast.Ident{Name: "x"}})
	if got != "x" {
		t.Fatalf("expected x, got %s", got)
	}
}
