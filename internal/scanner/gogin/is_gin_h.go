//ff:func feature=scan type=extract control=sequence
//ff:what CompositeLit이 gin.H 타입인지 검사한다
package gogin

import (
	"go/ast"
	"go/types"
)

func isGinH(comp *ast.CompositeLit, info *types.Info) bool {
	if comp.Type == nil {
		return false
	}

	// AST에서 gin.H 패턴 확인
	if isGinHSelector(comp.Type) {
		return true
	}

	// go/types로도 확인
	return isGinHMapType(comp, info)
}
