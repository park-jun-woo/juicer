//ff:func feature=scan type=test control=sequence topic=actix
//ff:what firstStringArgOfScopedCall — 지정 scoped 호출의 첫 문자열 인자 추출을 검증
package actix

import "testing"

func TestFirstStringArgOfScopedCall(t *testing.T) {
	src := []byte(`fn f() { web::scope("/api").service(x); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := firstStringArgOfScopedCall(root, src, "web::scope"); got != "/api" {
		t.Fatalf("got %q, want /api", got)
	}
	// Different scoped name -> not matched.
	if got := firstStringArgOfScopedCall(root, src, "web::resource"); got != "" {
		t.Fatalf("expected empty for unmatched name, got %q", got)
	}
}
