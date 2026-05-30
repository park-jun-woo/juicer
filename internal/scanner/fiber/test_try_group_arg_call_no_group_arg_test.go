//ff:func feature=scan type=test control=sequence
//ff:what TestTryGroupArgCall_NoGroupArg 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryGroupArgCall_NoGroupArg(t *testing.T) {

	call := parseCall(t, "doThing(x, y)")
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{},
		idx:     &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info:    newEmptyInfo(),
	}
	tryGroupArgCall(call, ctx)
}
