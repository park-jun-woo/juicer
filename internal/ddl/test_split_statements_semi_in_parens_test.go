//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitStatements_SemiInParens 테스트
package ddl

import "testing"

func TestSplitStatements_SemiInParens(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (id INT DEFAULT (1;2)); SELECT 1")
	if len(stmts) != 2 {
		t.Fatalf("expected 2, got %d: %v", len(stmts), stmts)
	}
}
