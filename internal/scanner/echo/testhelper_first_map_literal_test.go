//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what firstMapLiteral 테스트 헬퍼: types.Info.Types에서 첫 map 컴포지트 리터럴 조회
package echo

import (
	"go/ast"
	"go/types"
)

// firstMapLiteral returns the first composite literal whose type is a map.
func firstMapLiteral(info *types.Info) *ast.CompositeLit {
	for e := range info.Types {
		cl, ok := e.(*ast.CompositeLit)
		if !ok {
			continue
		}
		if _, isMap := cl.Type.(*ast.MapType); isMap {
			return cl
		}
	}
	return nil
}
