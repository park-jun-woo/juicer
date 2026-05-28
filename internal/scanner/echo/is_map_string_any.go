//ff:func feature=scan type=extract control=sequence
//ff:what CompositeLit이 map[string]any 타입인지 검사한다 (Echo에서 gin.H 대신 사용)
package echo

import (
	"go/ast"
	"go/types"
)

func isMapStringAny(comp *ast.CompositeLit, info *types.Info) bool {
	if comp.Type == nil {
		return false
	}

	// go/types로 확인
	tv, ok := info.Types[comp]
	if !ok {
		return false
	}
	mp, ok := tv.Type.Underlying().(*types.Map)
	if !ok {
		return false
	}
	return mp.Key().String() == "string"
}
