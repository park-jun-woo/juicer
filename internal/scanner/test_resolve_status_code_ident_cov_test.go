//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_IdentCov 테스트
package scanner

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveStatusCode_IdentCov(t *testing.T) {
	ident := &ast.Ident{Name: "statusCode"}
	c := types.NewConst(token.NoPos, nil, "statusCode", types.Typ[types.Int], constant.MakeInt64(404))
	info := &types.Info{
		Uses:  map[*ast.Ident]types.Object{ident: c},
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	got := resolveStatusCode(ident, info)
	if got != "404" {
		t.Fatalf("expected 404, got %s", got)
	}
}
