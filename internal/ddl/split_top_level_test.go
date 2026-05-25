package ddl

import "testing"

func TestSplitTopLevel_Basic(t *testing.T) {
	parts := splitTopLevel("a, b, c", ',')
	if len(parts) != 3 {
		t.Fatalf("expected 3, got %d", len(parts))
	}
}

func TestSplitTopLevel_WithParens(t *testing.T) {
	parts := splitTopLevel("a(1,2), b", ',')
	if len(parts) != 2 {
		t.Fatalf("expected 2, got %d: %v", len(parts), parts)
	}
}

func TestSplitTopLevel_Empty(t *testing.T) {
	parts := splitTopLevel("", ',')
	if len(parts) != 0 {
		t.Fatalf("expected 0, got %d", len(parts))
	}
}
