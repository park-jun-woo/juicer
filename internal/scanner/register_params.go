//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what registerParams 함수
package scanner

import (
	"go/ast"
)

// registerParams — 함수 파라미터 중 *gin.Engine / *gin.RouterGroup을 라우터로 등록
func registerParams(fn *ast.FuncDecl, ginAlias string, routers map[string]*routerInfo) {
	if fn.Type.Params == nil {
		return
	}
	for _, field := range fn.Type.Params.List {
		if !isGinRouterType(field.Type, ginAlias) {
			continue
		}
		for _, name := range field.Names {
			routers[name.Name] = &routerInfo{}
		}
	}
}
