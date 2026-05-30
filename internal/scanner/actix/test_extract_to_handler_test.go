//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractToHandler 테스트
package actix

import "testing"

func TestExtractToHandler(t *testing.T) {
	src := []byte(`fn f() { web::get().to(my_handler); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".to")
	if call == nil {
		t.Fatal(".to call not found")
	}
	if got := extractToHandler(call, src); got != "my_handler" {
		t.Fatalf("extractToHandler = %q, want my_handler", got)
	}
}
