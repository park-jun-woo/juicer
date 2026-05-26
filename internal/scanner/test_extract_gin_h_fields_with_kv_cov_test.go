//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGinHFields_WithKVCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestExtractGinHFields_WithKVCov(t *testing.T) {
	comp := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "message"},
				Value: &ast.Ident{Name: "val"},
			},
			&ast.Ident{Name: "skip"},
		},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	result := extractGinHFields(comp, info)
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
}
