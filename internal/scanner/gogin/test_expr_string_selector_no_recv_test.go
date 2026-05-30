//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_SelectorNoRecv 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_SelectorNoRecv(t *testing.T) {

	sel := &ast.SelectorExpr{X: &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}, Sel: &ast.Ident{Name: "M"}}
	if got := exprString(sel); got == "" {
		t.Error("expected non-empty selector string")
	}
}
