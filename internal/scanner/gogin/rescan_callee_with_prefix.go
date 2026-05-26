//ff:func feature=scan type=extract control=sequence
//ff:what 호출 대상 함수를 resolve하고, prefix를 전파하며 body를 재스캔하여 기존 endpoints를 업데이트한다
package gogin

import (
	"go/ast"
)

const maxRescanDepth = 2

func rescanCalleeWithPrefix(call *ast.CallExpr, argIdx int, prefix string, parent *routerInfo, ctx *groupArgCtx) {
	rescanCalleeWithPrefixDepth(call, argIdx, prefix, parent, ctx, 0)
}
