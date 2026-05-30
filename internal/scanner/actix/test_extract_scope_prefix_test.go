//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractScopePrefix 테스트
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
