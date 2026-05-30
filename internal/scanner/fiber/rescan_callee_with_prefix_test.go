//ff:func feature=scan type=test control=sequence
//ff:what rescanCalleeWithPrefix — depth 0 위임 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestRescanCalleeWithPrefix(t *testing.T) {
	call := parseCall(t, "registerRoutes(app)")
	ctx := &groupArgCtx{
		idx:  &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info: newEmptyInfo(),
	}
	// empty info -> resolveCallTarget invalid -> returns without panic.
	rescanCalleeWithPrefix(call, 0, "/api", &routerInfo{prefix: "/api"}, ctx)
}
