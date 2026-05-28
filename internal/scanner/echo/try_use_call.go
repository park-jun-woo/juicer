//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Use() 호출이면 해당 라우터의 미들웨어 목록에 추가한다
package echo

import (
	"go/ast"
)

func tryUseCall(call *ast.CallExpr, routers map[string]*routerInfo) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}
	// Echo는 Use()와 Pre() 모두 미들웨어 등록
	if sel.Sel.Name != "Use" && sel.Sel.Name != "Pre" {
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
