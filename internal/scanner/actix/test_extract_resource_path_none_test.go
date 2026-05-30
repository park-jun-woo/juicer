//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractResourcePath_None 테스트
package actix

import "testing"

func TestExtractResourcePath_None(t *testing.T) {
	src := []byte(`fn f() { web::scope("/p"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := extractResourcePath(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
