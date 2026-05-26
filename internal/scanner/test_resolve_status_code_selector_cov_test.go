//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_SelectorCov 테스트
package scanner

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveStatusCode_SelectorCov(t *testing.T) {
	selIdent := &ast.Ident{Name: "StatusOK"}
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "http"}, Sel: selIdent}
	c := types.NewConst(token.NoPos, nil, "StatusOK", types.Typ[types.Int], constant.MakeInt64(200))
	info := &types.Info{
		Uses:  map[*ast.Ident]types.Object{selIdent: c},
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	got := resolveStatusCode(sel, info)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}
