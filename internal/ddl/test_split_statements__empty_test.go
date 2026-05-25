//ff:func feature=ddl type=parse control=sequence
//ff:what TestSplitStatements_Empty 테스트
package ddl

import "testing"

func TestSplitStatements_Empty(t *testing.T) {
	stmts := splitStatements("")
	if len(stmts) != 0 {
		t.Fatalf("expected 0, got %d", len(stmts))
	}
}
