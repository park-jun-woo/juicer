//ff:func feature=scan type=test control=sequence
//ff:what TestRescanCalleeWithPrefixDepth_MaxDepth 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestRescanCalleeWithPrefixDepth_MaxDepth(t *testing.T) {
	call := parseCall(t, "registerRoutes(app)")
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}, info: newEmptyInfo()}

	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, maxRescanDepth)
}
