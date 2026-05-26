//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitTopLevel_WithParensCov 테스트
package ddl

import "testing"

func TestSplitTopLevel_WithParensCov(t *testing.T) {
	parts := splitTopLevel("a(1,2), b", ',')
	if len(parts) != 2 {
		t.Fatalf("expected 2, got %d: %v", len(parts), parts)
	}
}
