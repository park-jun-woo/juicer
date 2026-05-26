//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_SelectorExpr 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_SelectorExpr(t *testing.T) {
	got := exprString(&ast.SelectorExpr{
		X:   &ast.Ident{Name: "pkg"},
		Sel: &ast.Ident{Name: "Func"},
	})
	if got != "pkg.Func" {
		t.Fatalf("expected pkg.Func, got %s", got)
	}
}
