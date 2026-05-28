//ff:func feature=scan type=extract control=sequence
//ff:what 함수 선언의 argIdx번째 파라미터가 *fiber.App/*fiber.Group이면 해당 이름을 반환한다
package fiber

import (
	"go/ast"
	"go/types"
)

func fiberRouterParamAtIndex(fn *ast.FuncDecl, info *types.Info, argIdx int) string {
	if fn.Type.Params == nil {
		return ""
	}
	field, name := paramFieldAtIndex(fn.Type.Params, argIdx)
	if field == nil || name == "" {
		return ""
	}
	if info == nil {
		return ""
	}
	t := info.TypeOf(field.Type)
	if t != nil && isFiberRouterTypeInfo(t) {
		return name
	}
	return ""
}
