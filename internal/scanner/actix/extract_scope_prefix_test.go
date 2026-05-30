//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractScopePrefix — web::scope("...") 프리픽스 추출을 검증
package actix

import "testing"

func TestExtractScopePrefix(t *testing.T) {
	src := []byte(`fn f() { web::scope("/api/v1").service(x); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := extractScopePrefix(root, src); got != "/api/v1" {
		t.Fatalf("extractScopePrefix = %q, want /api/v1", got)
	}
}

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
