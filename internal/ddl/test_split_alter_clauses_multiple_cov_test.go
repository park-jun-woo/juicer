//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitAlterClauses_MultipleCov 테스트
package ddl

import "testing"

func TestSplitAlterClauses_MultipleCov(t *testing.T) {
	clauses := splitAlterClauses("ADD COLUMN name TEXT, DROP COLUMN age")
	if len(clauses) != 2 {
		t.Fatalf("expected 2, got %d: %v", len(clauses), clauses)
	}
}
