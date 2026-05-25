//ff:func feature=scan type=extract control=sequence
//ff:what TestUnquote_Invalid 테스트
package scanner

import "testing"

func TestUnquote_Invalid(t *testing.T) {
	got := unquote("hello")
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}
