//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseExtractor_Generic 테스트
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
