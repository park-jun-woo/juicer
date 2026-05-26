//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitStatements_SemicolonInParensCov 테스트
package ddl

import "testing"

func TestSplitStatements_SemicolonInParensCov(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (def TEXT DEFAULT ';')")
	if len(stmts) != 1 {
		t.Fatalf("expected 1, got %d", len(stmts))
	}
}
