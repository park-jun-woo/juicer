//ff:func feature=ddl type=parse control=sequence
//ff:what TestSplitTopLevel_Basic 테스트
package ddl

import "testing"

func TestSplitTopLevel_Basic(t *testing.T) {
	parts := splitTopLevel("a, b, c", ',')
	if len(parts) != 3 {
		t.Fatalf("expected 3, got %d", len(parts))
	}
}
