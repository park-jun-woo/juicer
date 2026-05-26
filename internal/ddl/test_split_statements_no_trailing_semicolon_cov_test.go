//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitStatements_NoTrailingSemicolonCov 테스트
package ddl

import "testing"

func TestSplitStatements_NoTrailingSemicolonCov(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (id INT)")
	if len(stmts) != 1 {
		t.Fatalf("expected 1, got %d", len(stmts))
	}
}
