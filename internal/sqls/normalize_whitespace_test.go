//ff:func feature=sql type=parse control=sequence
//ff:what TestNormalizeWhitespace_Spaces 테스트
package sqls

import "testing"

func TestNormalizeWhitespace_Spaces(t *testing.T) {
	got := normalizeWhitespace("  a   b   c  ")
	if got != "a b c" {
		t.Fatalf("expected 'a b c', got '%s'", got)
	}
}
