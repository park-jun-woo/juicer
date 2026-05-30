//ff:func feature=scan type=test control=sequence
//ff:what TestRescanCalleeWithPrefixDepth_InvalidTarget 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestRescanCalleeWithPrefixDepth_InvalidTarget(t *testing.T) {
	call := goginParseCall(t, "registerRoutes(r)")
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}, info: goginEmptyInfo()}
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, 0)
}
