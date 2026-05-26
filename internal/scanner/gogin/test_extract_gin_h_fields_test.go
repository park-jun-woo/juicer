//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractGinHFields 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestExtractGinHFields(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	comp := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.BasicLit{Kind: token.STRING, Value: `"message"`},
				Value: &ast.BasicLit{Kind: token.STRING, Value: `"ok"`},
			},
			&ast.KeyValueExpr{
				Key:   &ast.BasicLit{Kind: token.STRING, Value: `"count"`},
				Value: &ast.BasicLit{Kind: token.INT, Value: "42"},
			},
		},
	}

	fields := extractGinHFields(comp, info)
	if len(fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(fields))
	}
}
