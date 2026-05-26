//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGinHFields_Empty 테스트
package gogin

import (
	"go/ast"
	"go/token"
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

	// with elements
	comp2 := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.BasicLit{Kind: token.STRING, Value: `"msg"`},
				Value: &ast.BasicLit{Kind: token.STRING, Value: `"ok"`},
			},
			&ast.Ident{Name: "x"}, // non-KV element -> nil
		},
	}
	result2 := extractGinHFields(comp2, info)
	if len(result2) != 1 {
		t.Fatalf("expected 1, got %d", len(result2))
	}
}

