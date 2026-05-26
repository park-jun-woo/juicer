//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_SelectorCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprName_SelectorCov(t *testing.T) {
	got := exprName(&ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "Get"}})
	if got != "h.Get" {
		t.Fatalf("expected h.Get, got %s", got)
	}
}
