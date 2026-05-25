//ff:func feature=ddl type=parse control=sequence
//ff:what TestSplitStatements_NoTrailingSemicolon 테스트
package ddl

import "testing"

func TestSplitStatements_NoTrailingSemicolon(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (id INT)")
	if len(stmts) != 1 {
		t.Fatalf("expected 1, got %d", len(stmts))
	}
}
