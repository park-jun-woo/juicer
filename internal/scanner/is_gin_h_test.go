package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinH_NilType(t *testing.T) {
	comp := &ast.CompositeLit{}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if isGinH(comp, info) {
		t.Fatal("expected false for nil type")
	}
}

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
