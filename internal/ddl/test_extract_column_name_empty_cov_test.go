//ff:func feature=ddl type=test control=sequence
//ff:what TestExtractColumnName_EmptyCov 테스트
package ddl

import "testing"

func TestExtractColumnName_EmptyCov(t *testing.T) {
	got := extractColumnName("")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
