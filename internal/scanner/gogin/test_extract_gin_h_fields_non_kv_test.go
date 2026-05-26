//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractGinHFields_NonKV 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestExtractGinHFields_NonKV(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	comp := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.Ident{Name: "x"}, // not a KeyValueExpr
		},
	}

	fields := extractGinHFields(comp, info)
	if len(fields) != 0 {
		t.Errorf("expected 0 fields for non-KV, got %d", len(fields))
	}
}
