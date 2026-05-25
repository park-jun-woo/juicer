package ddl

import "testing"

func TestSplitStatements_Basic(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (id INT); CREATE TABLE b (id INT);")
	if len(stmts) != 2 {
		t.Fatalf("expected 2, got %d", len(stmts))
	}
}

func TestSplitStatements_Empty(t *testing.T) {
	stmts := splitStatements("")
	if len(stmts) != 0 {
		t.Fatalf("expected 0, got %d", len(stmts))
	}
}

func TestSplitStatements_NoTrailingSemicolon(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (id INT)")
	if len(stmts) != 1 {
		t.Fatalf("expected 1, got %d", len(stmts))
	}
}
