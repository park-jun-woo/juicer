//ff:func feature=ddl type=parse control=sequence
//ff:what TestSplitAlterClauses_Multiple 테스트
package ddl

import "testing"

func TestSplitAlterClauses_Multiple(t *testing.T) {
	clauses := splitAlterClauses("ADD COLUMN name TEXT, DROP COLUMN email")
	if len(clauses) != 2 {
		t.Fatalf("expected 2, got %d: %v", len(clauses), clauses)
	}
}
