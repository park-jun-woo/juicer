//ff:func feature=scan type=test control=sequence
//ff:what TestForwardRouterCalls_ExprStmtNonCall 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestForwardRouterCalls_ExprStmtNonCall(t *testing.T) {

	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.UnaryExpr{Op: token.ARROW, X: &ast.Ident{Name: "ch"}}},
	}
	ctx := &groupArgCtx{
		info: &types.Info{Uses: map[*ast.Ident]types.Object{}},
		fset: token.NewFileSet(),
		idx:  &funcIndex{},
	}
	forwardRouterCalls(stmts, "router", "/api", &routerInfo{}, nil, ctx, 0)
}
