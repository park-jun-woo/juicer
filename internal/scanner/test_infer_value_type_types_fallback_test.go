//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_TypesFallback 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_TypesFallback(t *testing.T) {
	// Create an expr that doesn't match any specific case but has Types info
	expr := &ast.ParenExpr{X: &ast.Ident{Name: "x"}}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{
			expr: {Type: types.Typ[types.Int]},
		},
	}
	got := inferValueType(expr, info)
	if got != "int" {
		t.Errorf("expected 'int', got %q", got)
	}
}
