//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractToHandler_ClosureOnly 테스트
package actix

import "testing"

func TestExtractToHandler_ClosureOnly(t *testing.T) {

	src := []byte(`fn f() { web::get().to(|| async { "x" }); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".to")
	if call == nil {
		t.Fatal(".to call not found")
	}
	if got := extractToHandler(call, src); got != "" {
		t.Fatalf("expected empty handler for closure, got %q", got)
	}
}
