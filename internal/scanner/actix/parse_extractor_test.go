//ff:func feature=scan type=test control=sequence topic=actix
//ff:what parseExtractor — 파라미터 extractor 파싱 분기를 검증
package actix

import "testing"

func TestParseExtractor_Generic(t *testing.T) {
	src := []byte(`fn f(body: web::Json<User>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	p := firstParam(root)
	if p == nil {
		t.Fatal("no parameter")
	}
	ext := parseExtractor(p, src)
	if ext == nil {
		t.Fatal("expected extractor info")
	}
	if ext.kind != "json" {
		t.Errorf("kind = %q, want json", ext.kind)
	}
	if ext.typeName != "User" {
		t.Errorf("typeName = %q, want User", ext.typeName)
	}
}

func TestParseExtractor_NonGeneric(t *testing.T) {
	// A primitive/type_identifier param type is not generic -> nil.
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

func TestParseExtractor_NoType(t *testing.T) {
	// A param node without a recognized type -> findParamType nil -> nil.
	// Use the self-parameter "&self" which has no type_identifier child.
	src := []byte(`impl S { fn f(&self) {} }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	// self_parameter is not a "parameter"; use the block to force the nil path.
	if ext := parseExtractor(root, src); ext != nil {
		t.Fatalf("expected nil, got %+v", ext)
	}
}
