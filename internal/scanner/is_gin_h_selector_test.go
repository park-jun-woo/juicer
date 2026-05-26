//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinHSelector_Valid 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinHSelector_Valid(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "H"}}
	if !isGinHSelector(sel) {
		t.Fatal("expected true")
	}
}
