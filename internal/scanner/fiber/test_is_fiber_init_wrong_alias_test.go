//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberInit_WrongAlias 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberInit_WrongAlias(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "New"}}
	if isFiberInit(sel, "fiber") {
		t.Fatal("expected false — wrong alias")
	}
}
