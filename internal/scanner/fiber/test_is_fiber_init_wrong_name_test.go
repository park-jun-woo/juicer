//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberInit_WrongName 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberInit_WrongName(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "other"}, Sel: &ast.Ident{Name: "New"}}
	if isFiberInit(sel, "fiber") {
		t.Fatal("expected false for wrong alias")
	}
	sel2 := &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "Config"}}
	if isFiberInit(sel2, "fiber") {
		t.Fatal("expected false for non-New method")
	}
}
