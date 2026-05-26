//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinInit_NotGin 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIsGinInit_NotGin(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "http"}, Sel: &ast.Ident{Name: "Default"}}
	if isGinInit(sel, "gin") {
		t.Fatal("expected false")
	}
}
