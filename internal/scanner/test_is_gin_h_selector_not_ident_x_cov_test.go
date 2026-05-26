//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinHSelector_NotIdentXCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinHSelector_NotIdentXCov(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "H"}}
	if isGinHSelector(sel) {
		t.Fatal("expected false")
	}
}
