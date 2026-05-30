//ff:func feature=scan type=test control=sequence
//ff:what TestWalkForGroupArgs_ExprStmtNonCall 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkForGroupArgs_ExprStmtNonCall(t *testing.T) {

	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.UnaryExpr{Op: token.ARROW, X: &ast.Ident{Name: "ch"}}},
	}
	ctx := &groupArgCtx{
		ginAlias: "gin",
		routers:  map[string]*routerInfo{},
		idx:      &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
	}
	walkForGroupArgs(stmts, ctx)
	_ = scanner.Endpoint{}
}
