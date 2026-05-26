//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_StarExprCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_StarExprCov(t *testing.T) {
	got := exprString(&ast.StarExpr{X: &ast.Ident{Name: "T"}})
	if got != "*T" {
		t.Fatalf("expected *T, got %s", got)
	}
}
