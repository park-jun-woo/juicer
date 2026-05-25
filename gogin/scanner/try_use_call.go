//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Use() 호출이면 해당 라우터의 미들웨어 목록에 추가한다
package scanner

import (
	"go/ast"
)

func tryUseCall(call *ast.CallExpr, routers map[string]*routerInfo) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Use" {
		return
	}
	recv := identName(sel.X)
	router, ok := routers[recv]
	if !ok {
		return
	}
	for _, arg := range call.Args {
		if name := exprName(arg); name != "" {
			router.middleware = append(router.middleware, name)
		}
	}
}

