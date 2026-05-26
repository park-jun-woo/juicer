//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinInit_NotIdentCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinInit_NotIdentCov(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "Default"}}
	if isGinInit(sel, "gin") {
		t.Fatal("expected false")
	}
}
