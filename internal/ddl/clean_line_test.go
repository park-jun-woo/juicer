//ff:func feature=ddl type=parse control=sequence
//ff:what TestCleanLine_WithComment 테스트
package ddl

import "testing"

func TestCleanLine_WithComment(t *testing.T) {
	// inline comment removed, surrounding whitespace trimmed
	got := cleanLine("  id INT NOT NULL -- primary key  ")
	if got != "id INT NOT NULL" {
		t.Fatalf("got %q", got)
	}

	// no comment: only whitespace trimmed
	got = cleanLine("  name TEXT  ")
	if got != "name TEXT" {
		t.Fatalf("no-comment line: got %q, want %q", got, "name TEXT")
	}
}
