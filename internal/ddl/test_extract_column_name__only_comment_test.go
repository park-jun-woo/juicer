//ff:func feature=ddl type=parse control=sequence
//ff:what TestExtractColumnName_OnlyComment 테스트
package ddl

import "testing"

func TestExtractColumnName_OnlyComment(t *testing.T) {
	got := extractColumnName("-- just a comment")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}
