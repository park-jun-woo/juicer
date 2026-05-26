//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_CompositeLitCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveExprType_CompositeLitCov(t *testing.T) {
	comp := &ast.CompositeLit{}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Defs:  make(map[*ast.Ident]types.Object),
		Types: map[ast.Expr]types.TypeAndValue{comp: {Type: types.Typ[types.Int]}},
	}
	resolveExprType(comp, info)
}
