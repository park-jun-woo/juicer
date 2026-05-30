//ff:func feature=scan type=test control=sequence
//ff:what TestWalkForGroupArgs_Empty 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestWalkForGroupArgs_Empty(t *testing.T) {
	ctx := &groupArgCtx{
		ginAlias:  "gin",
		routers:   map[string]*routerInfo{},
		fset:      token.NewFileSet(),
		endpoints: []scanner.Endpoint{},
		hmap:      map[int][]ast.Expr{},
		epIndex:   map[struct{ file string; line int }]int{},
	}
	// nil stmts should not panic
	walkForGroupArgs(nil, ctx)
}

func TestWalkForGroupArgs_ExprStmtNonCall(t *testing.T) {
	// ExprStmt whose X is not a CallExpr -> continue
	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.UnaryExpr{Op: token.ARROW, X: &ast.Ident{Name: "ch"}}},
	}
	ctx := &groupArgCtx{
		ginAlias: "gin",
		routers:    map[string]*routerInfo{},
		idx:        &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
	}
	walkForGroupArgs(stmts, ctx)
	_ = scanner.Endpoint{}
}
