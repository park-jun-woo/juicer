//ff:func feature=ddl type=test control=sequence
//ff:what TestExtractColumnName_Basic 테스트
package ddl

import "testing"

func TestExtractColumnName_Basic(t *testing.T) {
	got := extractColumnName("id INT PRIMARY KEY")
	if got != "id" {
		t.Fatalf("got %q", got)
	}

	// line comment without newline
	if extractColumnName("-- just a comment") != "" {
		t.Fatal("expected empty for comment-only line")
	}

	// line comment with newline then column def
	if extractColumnName("-- comment\nname TEXT") != "name" {
		t.Fatal("expected name after comment")
	}

	// empty/whitespace
	if extractColumnName("   ") != "" {
		t.Fatal("expected empty for whitespace")
	}
}
