//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what parseIntDefault 테스트
package fastapi

import "testing"

func TestParseIntDefault(t *testing.T) {
	if got := parseIntDefault("42", 0); got != 42 {
		t.Errorf("got %d, want 42", got)
	}
	if got := parseIntDefault("", 5); got != 5 {
		t.Errorf("empty: got %d, want 5", got)
	}
	if got := parseIntDefault("abc", 10); got != 10 {
		t.Errorf("abc: got %d, want 10", got)
	}
	if got := parseIntDefault("0", 99); got != 0 {
		t.Errorf("0: got %d, want 0", got)
	}
}
