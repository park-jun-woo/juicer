//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_StarExpr 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_StarExpr(t *testing.T) {
	got := exprString(&ast.StarExpr{X: &ast.Ident{Name: "int"}})
	if got != "*int" {
		t.Fatalf("expected *int, got %s", got)
	}
}
