//ff:func feature=scan type=extract control=sequence
//ff:what go/types에서 CompositeLit이 map[string]any인지 검사한다
package scanner

import (
	"go/ast"
	"go/types"
)

// isGinHMapType checks if a CompositeLit is a map[string]any via go/types.
func isGinHMapType(comp *ast.CompositeLit, info *types.Info) bool {
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
