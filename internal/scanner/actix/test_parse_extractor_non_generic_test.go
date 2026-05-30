//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseExtractor_NonGeneric 테스트
package actix

import "testing"

func TestParseExtractor_NonGeneric(t *testing.T) {

	src := []byte(`fn f(x: String) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	p := firstParam(root)
	if p == nil {
		t.Fatal("no parameter")
	}
	if ext := parseExtractor(p, src); ext != nil {
		t.Fatalf("expected nil for non-generic type, got %+v", ext)
	}
}
