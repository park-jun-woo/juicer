//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitAlterClauses_Single 테스트
package ddl

import "testing"

func TestSplitAlterClauses_Single(t *testing.T) {
	clauses := splitAlterClauses("ADD COLUMN name TEXT")
	if len(clauses) != 1 {
		t.Fatalf("expected 1, got %d", len(clauses))
	}

	// multiple clauses
	clauses = splitAlterClauses("ADD COLUMN name TEXT, DROP COLUMN age, ADD COLUMN email TEXT")
	if len(clauses) != 3 {
		t.Fatalf("expected 3, got %d: %v", len(clauses), clauses)
	}

	// continuation (non-ALTER keyword part)
	clauses = splitAlterClauses("something, else")
	if len(clauses) != 1 {
		t.Fatalf("expected 1 (continuation), got %d: %v", len(clauses), clauses)
	}

	// empty
	clauses = splitAlterClauses("")
	if len(clauses) != 0 {
		t.Fatalf("expected 0, got %d", len(clauses))
	}
}
