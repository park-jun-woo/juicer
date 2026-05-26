//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_TypesInfoCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_TypesInfoCov(t *testing.T) {
	ident := &ast.Ident{Name: "val"}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{
			ident: {Type: types.Typ[types.Int]},
		},
	}
	if got := inferValueType(ident, info); got != "int" {
		t.Fatalf("expected int, got %s", got)
	}
}
