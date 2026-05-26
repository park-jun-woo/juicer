//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitAlterClauses_ContinuationCov 테스트
package ddl

import "testing"

func TestSplitAlterClauses_ContinuationCov(t *testing.T) {
	// A non-clause part continues from previous
	clauses := splitAlterClauses("some prefix, ADD COLUMN name TEXT")
	if len(clauses) < 1 {
		t.Fatal("expected at least 1 clause")
	}
}
