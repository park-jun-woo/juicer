//ff:func feature=scan type=test control=iteration dimension=1
//ff:what forwardRouterCalls — 라우터 전달 호출 처리 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func bodyStmts(t *testing.T, src string) []ast.Stmt {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Body != nil {
			return fn.Body.List
		}
	}
	t.Fatal("no func body")
	return nil
}

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
	// empty info -> rescan resolves nothing and returns early; no panic.
	forwardRouterCalls(stmts, "app", "/api", &routerInfo{prefix: "/api"}, newEmptyInfo(), ctx, 0)
}

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
