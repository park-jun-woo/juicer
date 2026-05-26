//ff:func feature=ddl type=test control=sequence
//ff:what TestExtractColumnName_CommentOnly 테스트
package ddl

import "testing"

func TestExtractColumnName_CommentOnly(t *testing.T) {
	got := extractColumnName("-- just a comment")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
