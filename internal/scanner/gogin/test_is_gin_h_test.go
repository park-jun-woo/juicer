//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinH 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinH(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	t.Run("nil type", func(t *testing.T) {
		comp := &ast.CompositeLit{}
		if isGinH(comp, info) {
			t.Error("expected false for nil Type")
		}
	})

	t.Run("gin.H selector", func(t *testing.T) {
		comp := &ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "gin"},
				Sel: &ast.Ident{Name: "H"},
			},
		}
		if !isGinH(comp, info) {
			t.Error("expected true for gin.H")
		}
	})

	t.Run("non-gin selector", func(t *testing.T) {
		comp := &ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "other"},
				Sel: &ast.Ident{Name: "H"},
			},
		}
		if isGinH(comp, info) {
			t.Error("expected false for other.H")
		}
	})

	t.Run("map via TypesInfo", func(t *testing.T) {
		comp := &ast.CompositeLit{
			Type: &ast.Ident{Name: "MyMap"},
		}
		mapType := types.NewMap(types.Typ[types.String], types.NewInterfaceType(nil, nil))
		infoWithTypes := &types.Info{
			Types: map[ast.Expr]types.TypeAndValue{
				comp: {Type: mapType},
			},
		}
		if !isGinH(comp, infoWithTypes) {
			t.Error("expected true for map[string]any via TypesInfo")
		}
	})

	t.Run("non-selector type", func(t *testing.T) {
		comp := &ast.CompositeLit{
			Type: &ast.Ident{Name: "MyMap"},
		}
		if isGinH(comp, info) {
			t.Error("expected false for non-selector")
		}
	})
}
