//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinHSelector_NonGin 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIsGinHSelector_NonGin(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "foo"}, Sel: &ast.Ident{Name: "H"}}
	if isGinHSelector(sel) {
		t.Fatal("expected false")
	}
}
