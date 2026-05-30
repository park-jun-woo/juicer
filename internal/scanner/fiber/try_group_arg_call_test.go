//ff:func feature=scan type=test control=iteration dimension=1
//ff:what tryGroupArgCall — 그룹 인자 호출 감지 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryGroupArgCall_GroupArg(t *testing.T) {
	// register(api) where api is a known router -> extractGroupArgPrefix matches,
	// then rescanCalleeWithPrefix runs (empty info -> early return, no panic).
	call := parseCall(t, "register(api)")
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{"api": {prefix: "/api"}},
		idx:     &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info:    newEmptyInfo(),
	}
	tryGroupArgCall(call, ctx)
}

func TestTryGroupArgCall_NoGroupArg(t *testing.T) {
	// no arg is a router -> all continue
	call := parseCall(t, "doThing(x, y)")
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{},
		idx:     &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info:    newEmptyInfo(),
	}
	tryGroupArgCall(call, ctx)
}
