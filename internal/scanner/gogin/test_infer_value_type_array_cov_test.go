//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_ArrayCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_ArrayCov(t *testing.T) {
	comp := &ast.CompositeLit{Type: &ast.ArrayType{Elt: &ast.Ident{Name: "int"}}}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if got := inferValueType(comp, info); got != "array" {
		t.Fatalf("expected array, got %s", got)
	}
}
