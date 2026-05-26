//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_SelectorNoRecvCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprName_SelectorNoRecvCov(t *testing.T) {
	got := exprName(&ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "Get"}})
	if got != "Get" {
		t.Fatalf("expected Get, got %s", got)
	}
}
