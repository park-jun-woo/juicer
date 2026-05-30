//ff:func feature=scan type=test control=sequence
//ff:what TestForwardRouterCalls_ForwardsParam 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestForwardRouterCalls_ForwardsParam(t *testing.T) {
	src := `package m
func Setup(app int, ch chan int) {
	registerRoutes(app)     // forwards "app" at index 0
	doSomething(x)          // no app -> skipped
	y := 1                  // not an ExprStmt -> skipped
	<-ch                    // ExprStmt whose X is not a CallExpr -> skipped
}
`
	stmts := bodyStmts(t, src)
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}}

	forwardRouterCalls(stmts, "app", "/api", &routerInfo{prefix: "/api"}, newEmptyInfo(), ctx, 0)
}
