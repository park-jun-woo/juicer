//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_SelectorCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_SelectorCov(t *testing.T) {
	got := exprString(&ast.SelectorExpr{X: &ast.Ident{Name: "pkg"}, Sel: &ast.Ident{Name: "Func"}})
	if got != "pkg.Func" {
		t.Fatalf("expected pkg.Func, got %s", got)
	}
}
