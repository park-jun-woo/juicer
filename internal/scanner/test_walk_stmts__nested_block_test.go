//ff:func feature=scan type=extract control=sequence
//ff:what TestWalkStmts_NestedBlock 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkStmts_NestedBlock(t *testing.T) {
	stmts := []ast.Stmt{
		&ast.BlockStmt{List: []ast.Stmt{}},
	}
	var out []Endpoint
	walkStmts(stmts, "gin", "test.go", token.NewFileSet(), map[string]*routerInfo{}, &out)
	if len(out) != 0 {
		t.Fatal("expected empty")
	}
}
