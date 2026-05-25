package sqls

import "testing"

func TestNormalizeWhitespace_Spaces(t *testing.T) {
	got := normalizeWhitespace("  a   b   c  ")
	if got != "a b c" {
		t.Fatalf("expected 'a b c', got '%s'", got)
	}
}
