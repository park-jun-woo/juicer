//ff:func feature=scan type=test control=sequence
//ff:what TestForwardRouterCalls_NoForwardingCalls 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestForwardRouterCalls_NoForwardingCalls(t *testing.T) {
	src := `package m
func Setup(app int) {
	z := compute()
}
`
	stmts := bodyStmts(t, src)
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}}
	forwardRouterCalls(stmts, "app", "", nil, newEmptyInfo(), ctx, 0)
}
