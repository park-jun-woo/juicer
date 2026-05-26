//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitStatements_Basic 테스트
package ddl

import "testing"

func TestSplitStatements_Basic(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (id INT); CREATE TABLE b (id INT);")
	if len(stmts) != 2 {
		t.Fatalf("expected 2, got %d", len(stmts))
	}
}
