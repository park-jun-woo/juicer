//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinInit_Default 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinInit_Default(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Default"}}
	if !isGinInit(sel, "gin") {
		t.Fatal("expected true")
	}
}
