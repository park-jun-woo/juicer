//ff:func feature=scan type=test control=sequence
//ff:what TestResolveUsesConst_FoundCov 테스트
package gogin

import (
	"go/ast"
	"go/constant"
	"go/types"
	"testing"
)

func TestResolveUsesConst_FoundCov(t *testing.T) {
	ident := &ast.Ident{Name: "StatusOK"}
	c := types.NewConst(0, nil, "StatusOK", types.Typ[types.Int], constant.MakeInt64(200))
	info := &types.Info{Uses: map[*ast.Ident]types.Object{ident: c}}
	got := resolveUsesConst(info, ident)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}
