//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberInit_New 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberInit_New(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "New"}}
	if !isFiberInit(sel, "fiber") {
		t.Fatal("expected true")
	}
}
