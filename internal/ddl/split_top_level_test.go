//ff:func feature=ddl type=test control=sequence
//ff:what TestSplitTopLevel_Basic 테스트
package ddl

import "testing"

func TestSplitTopLevel_Basic(t *testing.T) {
	parts := splitTopLevel("a, b, c", ',')
	if len(parts) != 3 {
		t.Fatalf("expected 3, got %d", len(parts))
	}

	// comma inside parentheses should not split
	parts = splitTopLevel("a(x,y), b", ',')
	if len(parts) != 2 {
		t.Fatalf("expected 2 (paren), got %d: %v", len(parts), parts)
	}

	// empty
	parts = splitTopLevel("", ',')
	if len(parts) != 0 {
		t.Fatalf("expected 0, got %d", len(parts))
	}
}
