//ff:func feature=scan type=extract control=sequence
//ff:what TestWalkStmts_ExprStmtNonCall 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestWalkStmts_ExprStmtNonCall(t *testing.T) {
	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "x"}},
	}
	var out []scanner.Endpoint
	walkStmts(stmts, "gin", "test.go", token.NewFileSet(), map[string]*routerInfo{}, &out, map[int][]ast.Expr{})
	if len(out) != 0 {
		t.Fatal("expected empty")
	}
}
