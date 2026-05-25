//ff:func feature=scan type=extract control=sequence
//ff:what TestWalkStmts_NonCallExprStmt 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkStmts_NonCallExprStmt(t *testing.T) {
	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "x"}},
	}
	fset := token.NewFileSet()
	routers := make(map[string]*routerInfo)
	var out []Endpoint
	walkStmts(stmts, "gin", "test.go", fset, routers, &out)
	if len(out) != 0 {
		t.Errorf("expected 0, got %d", len(out))
	}
}
