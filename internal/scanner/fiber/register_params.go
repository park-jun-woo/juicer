//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 함수 파라미터 중 *fiber.App / *fiber.Group을 라우터로 등록한다
package fiber

import (
	"go/ast"
)

// registerParams — 함수 파라미터 중 *fiber.App / *fiber.Group을 라우터로 등록
func registerParams(fn *ast.FuncDecl, fiberAlias string, routers map[string]*routerInfo) {
	if fn.Type.Params == nil {
		return
	}
	for _, field := range fn.Type.Params.List {
		if !isFiberRouterType(field.Type, fiberAlias) {
			continue
		}
		for _, name := range field.Names {
			routers[name.Name] = &routerInfo{}
		}
	}
}
