//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseExtractor_NoType 테스트
package actix

import "testing"

func TestParseExtractor_NoType(t *testing.T) {

	src := []byte(`impl S { fn f(&self) {} }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}

	if ext := parseExtractor(root, src); ext != nil {
		t.Fatalf("expected nil, got %+v", ext)
	}
}
