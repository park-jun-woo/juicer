//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWebMethodFromCall_NoScopedIdentifier 테스트
package actix

import "testing"

func TestWebMethodFromCall_NoScopedIdentifier(t *testing.T) {

	src := []byte(`fn f() { x.foo(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".foo")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
