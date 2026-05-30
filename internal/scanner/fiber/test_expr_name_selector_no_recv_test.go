//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_SelectorNoRecv 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestExprName_SelectorNoRecv(t *testing.T) {

	sel := &ast.SelectorExpr{
		X:   &ast.CallExpr{Fun: ast.NewIdent("f")},
		Sel: ast.NewIdent("Method"),
	}
	if got := exprName(sel); got != "Method" {
		t.Errorf("selector no-recv: %q, want Method", got)
	}
}
