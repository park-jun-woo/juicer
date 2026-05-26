//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 함수 호출의 인자에서 router.Group("prefix") 패턴을 감지하고, 대상 함수의 라우트에 prefix를 전파한다
package gogin

import (
	"go/ast"
)

func tryGroupArgCall(call *ast.CallExpr, ctx *groupArgCtx) {
	for argIdx, arg := range call.Args {
		prefix, parent, ok := extractGroupArgPrefix(arg, ctx)
		if !ok {
			continue
		}
		rescanCalleeWithPrefix(call, argIdx, prefix, parent, ctx)
	}
}
