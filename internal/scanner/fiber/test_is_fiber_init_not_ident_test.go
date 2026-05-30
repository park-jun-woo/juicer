//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberInit_NotIdent 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberInit_NotIdent(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.CallExpr{}, Sel: &ast.Ident{Name: "New"}}
	if isFiberInit(sel, "fiber") {
		t.Fatal("expected false when X is not an ident")
	}
}
