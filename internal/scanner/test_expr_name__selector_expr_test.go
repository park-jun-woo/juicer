//ff:func feature=scan type=extract control=sequence
//ff:what TestExprName_SelectorExpr 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprName_SelectorExpr(t *testing.T) {
	got := exprName(&ast.SelectorExpr{
		X:   &ast.Ident{Name: "h"},
		Sel: &ast.Ident{Name: "Create"},
	})
	if got != "h.Create" {
		t.Fatalf("expected h.Create, got %s", got)
	}
}
