package sqls

import (
	"go/ast"
	"testing"
)

func TestRefineCRUDFromAST_Nil(t *testing.T) {
	got := refineCRUDFromAST(nil)
	if got != "EXEC" {
		t.Fatalf("expected EXEC, got %s", got)
	}
}

func TestRefineCRUDFromAST_Empty(t *testing.T) {
	got := refineCRUDFromAST(&ast.BlockStmt{})
	if got != "EXEC" {
		t.Fatalf("expected EXEC, got %s", got)
	}
}
