//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinInit_NonIdent 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinInit_NonIdent(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.BasicLit{Value: "x"}, Sel: &ast.Ident{Name: "Default"}}
	if isGinInit(sel, "gin") {
		t.Fatal("expected false")
	}
}
