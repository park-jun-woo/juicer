//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what types.Info 기반으로 함수 타입에서 *gin.Context 파라미터의 이름을 반환한다
package scanner

import (
	"go/ast"
	"go/types"
)

func ginCtxParamNameInfo(ft *ast.FuncType, info *types.Info) string {
	if ft.Params == nil || info == nil {
		return ginCtxParamName(ft)
	}
	for _, field := range ft.Params.List {
		t := info.TypeOf(field.Type)
		if t != nil && isGinContextTypeInfo(t) {
			if len(field.Names) > 0 {
				return field.Names[0].Name
			}
			return "_"
		}
	}
	// types.Info로 판별 실패 시 AST 기반 fallback
	return ginCtxParamName(ft)
}
