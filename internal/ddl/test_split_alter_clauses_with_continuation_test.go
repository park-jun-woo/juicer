//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitAlterClauses_WithContinuation 테스트
package ddl

import "testing"

func TestSplitAlterClauses_WithContinuation(t *testing.T) {
	clauses := splitAlterClauses("some prefix, ADD COLUMN name TEXT, DROP COLUMN age")
	if len(clauses) < 2 {
		t.Fatalf("expected at least 2, got %d: %v", len(clauses), clauses)
	}
}
