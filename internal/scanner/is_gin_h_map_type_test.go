package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinHMapType_NoTypesEntry(t *testing.T) {
	comp := &ast.CompositeLit{Type: &ast.Ident{Name: "X"}}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if isGinHMapType(comp, info) {
		t.Fatal("expected false")
	}
}
