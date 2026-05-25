package ddl

import "testing"

func TestSplitAlterClauses_Single(t *testing.T) {
	clauses := splitAlterClauses("ADD COLUMN name TEXT")
	if len(clauses) != 1 {
		t.Fatalf("expected 1, got %d", len(clauses))
	}
}

func TestSplitAlterClauses_Multiple(t *testing.T) {
	clauses := splitAlterClauses("ADD COLUMN name TEXT, DROP COLUMN email")
	if len(clauses) != 2 {
		t.Fatalf("expected 2, got %d: %v", len(clauses), clauses)
	}
}

func TestSplitAlterClauses_Empty(t *testing.T) {
	clauses := splitAlterClauses("")
	if len(clauses) != 0 {
		t.Fatalf("expected 0, got %d", len(clauses))
	}
}
