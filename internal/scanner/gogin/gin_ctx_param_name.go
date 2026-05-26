//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 함수 타입에서 *gin.Context 파라미터의 이름을 반환한다
package gogin

import (
	"go/ast"
)

func ginCtxParamName(ft *ast.FuncType) string {
	if ft.Params == nil {
		return ""
	}
	for _, field := range ft.Params.List {
		if isGinContextType(field.Type) {
			if len(field.Names) > 0 {
				return field.Names[0].Name
			}
			return "_"
		}
	}
	return ""
}

