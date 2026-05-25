package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkStmts_Empty(t *testing.T) {
	var out []Endpoint
	walkStmts(nil, "gin", "test.go", token.NewFileSet(), map[string]*routerInfo{}, &out)
	if len(out) != 0 {
		t.Fatal("expected empty")
	}
}

func TestWalkStmts_ExprStmtNonCall(t *testing.T) {
	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "x"}},
	}
	var out []Endpoint
	walkStmts(stmts, "gin", "test.go", token.NewFileSet(), map[string]*routerInfo{}, &out)
	if len(out) != 0 {
		t.Fatal("expected empty")
	}
}

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
