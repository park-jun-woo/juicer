//ff:func feature=scan type=extract control=sequence
//ff:what funcIndex에서 위치 기반으로 함수 선언과 TypesInfo를 찾는다
package fiber

import (
	"go/ast"
	"go/token"
	"go/types"
)

func lookupFunc(pos token.Pos, idx *funcIndex) (*ast.FuncDecl, *types.Info) {
	fn := idx.byPos[pos]
	if fn == nil {
		return nil, nil
	}
	return fn, idx.info[pos]
}
