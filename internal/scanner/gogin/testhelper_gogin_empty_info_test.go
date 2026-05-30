//ff:func feature=scan type=test control=sequence
//ff:what goginEmptyInfo 테스트 헬퍼
package gogin

import (
	"go/ast"
	"go/types"
)

func goginEmptyInfo() *types.Info {
	return &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
}
