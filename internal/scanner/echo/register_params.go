//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 함수 파라미터 중 *echo.Echo / *echo.Group을 라우터로 등록한다
package echo

import (
	"go/ast"
)

func registerParams(fn *ast.FuncDecl, echoAlias string, routers map[string]*routerInfo) {
	if fn.Type.Params == nil {
		return
	}
	for _, field := range fn.Type.Params.List {
		if !isEchoRouterType(field.Type, echoAlias) {
			continue
		}
		for _, name := range field.Names {
			routers[name.Name] = &routerInfo{}
		}
	}
}
