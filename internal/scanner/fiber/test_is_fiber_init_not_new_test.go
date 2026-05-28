//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberInit_NotNew 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberInit_NotNew(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "Default"}}
	if isFiberInit(sel, "fiber") {
		t.Fatal("expected false — Fiber only has New()")
	}
}
