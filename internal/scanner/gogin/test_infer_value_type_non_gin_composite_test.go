//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_NonGinComposite 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_NonGinComposite(t *testing.T) {

	lit := &ast.CompositeLit{Type: &ast.ArrayType{Elt: &ast.Ident{Name: "int"}}}
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{}}
	if got := inferValueType(lit, info); got != "array" {
		t.Fatalf("expected array, got %q", got)
	}
}
