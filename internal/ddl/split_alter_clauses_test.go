//ff:func feature=ddl type=parse control=sequence
//ff:what TestSplitAlterClauses_Single 테스트
package ddl

import "testing"

func TestSplitAlterClauses_Single(t *testing.T) {
	clauses := splitAlterClauses("ADD COLUMN name TEXT")
	if len(clauses) != 1 {
		t.Fatalf("expected 1, got %d", len(clauses))
	}
}
