//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestHandlerOrAnon — 핸들러명이 있으면 그대로, 없으면 익명 id를 반환하는지 검증
package actix

import "testing"

func TestHandlerOrAnon(t *testing.T) {
	if got := handlerOrAnon("foo", "GET", "/x"); got != "foo" {
		t.Errorf("named: got %q want foo", got)
	}
	if got := handlerOrAnon("", "GET", "/x"); got != "anon:GET:/x" {
		t.Errorf("empty: got %q want anon:GET:/x", got)
	}
	// distinct method/path yield distinct ids (no false dedup collision)
	if handlerOrAnon("", "GET", "/x") == handlerOrAnon("", "POST", "/x") {
		t.Error("anon ids for different methods must differ")
	}
}
