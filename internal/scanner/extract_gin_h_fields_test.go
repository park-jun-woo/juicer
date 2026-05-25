package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestExtractGinHFields_Empty(t *testing.T) {
	comp := &ast.CompositeLit{}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	result := extractGinHFields(comp, info)
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
