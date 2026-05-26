//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_CompositeLitArray 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_CompositeLitArray(t *testing.T) {
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	expr := &ast.CompositeLit{
		Type: &ast.ArrayType{Elt: &ast.Ident{Name: "int"}},
	}
	got := inferValueType(expr, info)
	if got != "array" {
		t.Errorf("expected 'array', got %q", got)
	}
}
