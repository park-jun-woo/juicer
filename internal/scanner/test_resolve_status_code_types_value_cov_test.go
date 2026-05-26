//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_TypesValueCov 테스트
package scanner

import (
	"go/ast"
	"go/constant"
	"go/types"
	"testing"
)

func TestResolveStatusCode_TypesValueCov(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
		Types: map[ast.Expr]types.TypeAndValue{
			expr: {Value: constant.MakeInt64(201)},
		},
	}
	got := resolveStatusCode(expr, info)
	if got != "201" {
		t.Fatalf("expected 201, got %s", got)
	}
}
