//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_SliceExpr 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_SliceExpr(t *testing.T) {
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	expr := &ast.SliceExpr{X: &ast.Ident{Name: "arr"}}
	got := inferValueType(expr, info)
	if got != "array" {
		t.Errorf("expected 'array', got %q", got)
	}
}
