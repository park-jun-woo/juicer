//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestUnquote_Round5 테스트
package echo

import "testing"

func TestUnquote_Round5(t *testing.T) {
	if unquote(`"hello"`) != "hello" {
		t.Fatal("double-quote")
	}
	if unquote("`raw`") != "raw" {
		t.Fatal("raw")
	}
	if unquote("plain") != "plain" {
		t.Fatal("plain")
	}
}
