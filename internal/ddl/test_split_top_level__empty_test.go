//ff:func feature=ddl type=parse control=sequence
//ff:what TestSplitTopLevel_Empty 테스트
package ddl

import "testing"

func TestSplitTopLevel_Empty(t *testing.T) {
	parts := splitTopLevel("", ',')
	if len(parts) != 0 {
		t.Fatalf("expected 0, got %d", len(parts))
	}
}
