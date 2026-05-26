//ff:func feature=scan type=extract control=sequence
//ff:what TestUnquote_Backtick 테스트
package gogin

import "testing"

func TestUnquote_Backtick(t *testing.T) {
	got := unquote("`hello`")
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}
