//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractToHandler_NoArgs 테스트
package actix

import "testing"

func TestExtractToHandler_NoArgs(t *testing.T) {

	src := []byte(`fn f() { web::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if call == nil {
		t.Fatal("get call not found")
	}
	if got := extractToHandler(call, src); got != "" {
		t.Fatalf("expected empty handler, got %q", got)
	}
}
