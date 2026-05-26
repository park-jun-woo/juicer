//ff:func feature=scan type=test control=sequence
//ff:what TestForwardRouterCalls_SkipsNonCallExpr 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestForwardRouterCalls_SkipsNonCallExpr(t *testing.T) {
	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "x"}},
	}
	ctx := &groupArgCtx{
		info: &types.Info{Uses: map[*ast.Ident]types.Object{}},
		fset: token.NewFileSet(),
		idx:  &funcIndex{},
	}
	parent := &routerInfo{prefix: "/api"}
	forwardRouterCalls(stmts, "router", "/api", parent, nil, ctx, 0)
}
