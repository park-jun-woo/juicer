//ff:func feature=scan type=test control=sequence
//ff:what unquote — 따옴표 제거 테스트
package fiber

import "testing"

func TestUnquote(t *testing.T) {
	// valid double-quoted
	if got := unquote(`"hello"`); got != "hello" {
		t.Errorf("double quote: %q", got)
	}
	// valid backtick
	if got := unquote("`raw`"); got != "raw" {
		t.Errorf("backtick: %q", got)
	}
	// invalid for strconv.Unquote -> Trim fallback
	if got := unquote(`"unterminated`); got != "unterminated" {
		t.Errorf("fallback trim: %q", got)
	}
}
