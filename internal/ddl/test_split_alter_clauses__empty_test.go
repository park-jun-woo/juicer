//ff:func feature=ddl type=parse control=sequence
//ff:what TestSplitAlterClauses_Empty 테스트
package ddl

import "testing"

func TestSplitAlterClauses_Empty(t *testing.T) {
	clauses := splitAlterClauses("")
	if len(clauses) != 0 {
		t.Fatalf("expected 0, got %d", len(clauses))
	}
}
