//ff:func feature=ddl type=parse control=sequence
//ff:what TestExtractColumnName_Empty 테스트
package ddl

import "testing"

func TestExtractColumnName_Empty(t *testing.T) {
	got := extractColumnName("")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}
