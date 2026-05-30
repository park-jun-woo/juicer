//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWebMethodFromCall_UnknownWebMethod 테스트
package actix

import "testing"

func TestWebMethodFromCall_UnknownWebMethod(t *testing.T) {

	src := []byte(`fn f() { web::unknownthing(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::unknownthing")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "" {
		t.Fatalf("expected empty for unknown web method, got %q", got)
	}
}
