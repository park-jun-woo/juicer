//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitStatements_Basic 테스트
package ddl

import "testing"

func TestSplitStatements_Basic(t *testing.T) {
	stmts := splitStatements("CREATE TABLE a (id INT); CREATE TABLE b (id INT);")
	if len(stmts) != 2 {
		t.Fatalf("expected 2, got %d", len(stmts))
	}

	// semicolon inside parentheses (should not split)
	stmts = splitStatements("CREATE TABLE a (check(x;y)); CREATE TABLE b (id INT)")
	if len(stmts) != 2 {
		t.Fatalf("expected 2 (semicolon in parens), got %d: %v", len(stmts), stmts)
	}

	// trailing statement without semicolon
	stmts = splitStatements("CREATE TABLE a (id INT)")
	if len(stmts) != 1 {
		t.Fatalf("expected 1 (no trailing semicolon), got %d", len(stmts))
	}
}
