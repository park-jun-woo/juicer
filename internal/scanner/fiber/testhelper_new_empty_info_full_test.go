//ff:func feature=scan type=test control=sequence
//ff:what newEmptyInfoFull 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/types"
)

func newEmptyInfoFull() *types.Info {
	return &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
}
