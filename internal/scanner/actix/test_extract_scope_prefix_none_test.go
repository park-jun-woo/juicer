//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractScopePrefix_None 테스트
package actix

import "testing"

func TestExtractScopePrefix_None(t *testing.T) {
	src := []byte(`fn f() { web::resource("/x"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := extractScopePrefix(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
