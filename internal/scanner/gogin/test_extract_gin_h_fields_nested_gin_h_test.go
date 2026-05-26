//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractGinHFields_NestedGinH 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestExtractGinHFields_NestedGinH(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	nested := &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "H"},
		},
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.BasicLit{Kind: token.STRING, Value: `"inner"`},
				Value: &ast.BasicLit{Kind: token.STRING, Value: `"val"`},
			},
		},
	}

	comp := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.BasicLit{Kind: token.STRING, Value: `"data"`},
				Value: nested,
			},
		},
	}

	fields := extractGinHFields(comp, info)
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	if fields[0].Type != "object" {
		t.Errorf("expected type 'object' for nested gin.H, got %q", fields[0].Type)
	}
	if len(fields[0].Fields) != 1 {
		t.Errorf("expected 1 nested field, got %d", len(fields[0].Fields))
	}
}
