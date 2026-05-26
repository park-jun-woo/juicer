//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinH_NonGinSelector 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinH_NonGinSelector(t *testing.T) {
	comp := &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "foo"},
			Sel: &ast.Ident{Name: "Bar"},
		},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if isGinH(comp, info) {
		t.Fatal("expected false for non-gin type")
	}
}
