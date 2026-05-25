//ff:func feature=ddl type=parse control=sequence
//ff:what TestExtractColumnName_WithComment 테스트
package ddl

import "testing"

func TestExtractColumnName_WithComment(t *testing.T) {
	got := extractColumnName("-- comment\nname TEXT")
	if got != "name" {
		t.Fatalf("got %q", got)
	}
}
