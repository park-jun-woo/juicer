//ff:func feature=ddl type=test control=sequence
//ff:what TestExtractColumnName_CommentThenCol 테스트
package ddl

import "testing"

func TestExtractColumnName_CommentThenCol(t *testing.T) {
	got := extractColumnName("-- comment\nname TEXT")
	if got != "name" {
		t.Fatalf("expected 'name', got %q", got)
	}
}
