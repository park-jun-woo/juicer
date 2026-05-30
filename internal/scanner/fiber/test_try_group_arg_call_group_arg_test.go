//ff:func feature=scan type=test control=sequence
//ff:what TestTryGroupArgCall_GroupArg 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryGroupArgCall_GroupArg(t *testing.T) {

	call := parseCall(t, "register(api)")
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{"api": {prefix: "/api"}},
		idx:     &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info:    newEmptyInfo(),
	}
	tryGroupArgCall(call, ctx)
}
