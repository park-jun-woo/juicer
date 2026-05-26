//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinH_GinHSelector 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinH_GinHSelector(t *testing.T) {
	comp := &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "H"},
		},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if !isGinH(comp, info) {
		t.Fatal("expected true for gin.H")
	}
}
