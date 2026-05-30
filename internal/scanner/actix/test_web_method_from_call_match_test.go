//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWebMethodFromCall_Match 테스트
package actix

import "testing"

func TestWebMethodFromCall_Match(t *testing.T) {
	src := []byte(`fn f() { web::post(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::post")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "POST" {
		t.Fatalf("got %q, want POST", got)
	}
}
