//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractGinHFields_IdentKey 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestExtractGinHFields_IdentKey(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	comp := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "key"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: `"val"`},
			},
		},
	}

	fields := extractGinHFields(comp, info)
	if len(fields) != 1 || fields[0].Name != "key" {
		t.Errorf("expected field 'key', got %v", fields)
	}
}
