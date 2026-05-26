//ff:func feature=ddl type=test control=sequence
//ff:what TestExtractColumnName_CommentWithColumnCov 테스트
package ddl

import "testing"

func TestExtractColumnName_CommentWithColumnCov(t *testing.T) {
	got := extractColumnName("-- comment\nname TEXT")
	if got != "name" {
		t.Fatalf("expected name, got %q", got)
	}
}
