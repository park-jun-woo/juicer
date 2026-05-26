//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_IdentUsesCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveExprType_IdentUsesCov(t *testing.T) {
	ident := &ast.Ident{Name: "req"}
	obj := types.NewVar(0, nil, "req", types.Typ[types.String])
	info := &types.Info{
		Uses:  map[*ast.Ident]types.Object{ident: obj},
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	resolveExprType(ident, info)
}
