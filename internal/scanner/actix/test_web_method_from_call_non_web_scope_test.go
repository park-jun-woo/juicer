//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWebMethodFromCall_NonWebScope 테스트
package actix

import "testing"

func TestWebMethodFromCall_NonWebScope(t *testing.T) {

	src := []byte(`fn f() { other::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "" {
		t.Fatalf("expected empty for non-web scope, got %q", got)
	}
}
