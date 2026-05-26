//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_UnaryExprCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExprString_UnaryExprCov(t *testing.T) {
	got := exprString(&ast.UnaryExpr{Op: token.AND, X: &ast.Ident{Name: "x"}})
	if got != "x" {
		t.Fatalf("expected x, got %s", got)
	}
}
