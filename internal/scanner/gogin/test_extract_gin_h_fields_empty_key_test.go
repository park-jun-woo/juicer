//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractGinHFields_EmptyKey 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestExtractGinHFields_EmptyKey(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	comp := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.CompositeLit{},
				Value: &ast.BasicLit{Kind: token.STRING, Value: `"val"`},
			},
		},
	}

	fields := extractGinHFields(comp, info)
	if len(fields) != 0 {
		t.Errorf("expected 0 fields for empty key, got %d", len(fields))
	}
}
