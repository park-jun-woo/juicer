//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestRescanCalleeWithPrefix_Wrapper_Round5 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestRescanCalleeWithPrefix_Wrapper_Round5(t *testing.T) {

	file, info := checkSrc(t, `package m
func helper() int { return 1 }
var _ = helper()
`)
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			call = c
		}
		return true
	})
	ctx := emptyGroupCtx()
	ctx.info = info

	rescanCalleeWithPrefix(call, 0, "/p", &routerInfo{}, ctx)
}
