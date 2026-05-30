//ff:func feature=scan type=test control=sequence
//ff:what newEmptyInfo 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/types"
)

func newEmptyInfo() *types.Info {
	return &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
}
