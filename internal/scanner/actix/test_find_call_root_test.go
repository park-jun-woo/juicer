//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindCallRoot 테스트
package actix

import "testing"

func TestFindCallRoot(t *testing.T) {
	src := []byte(`fn f() { web::scope("/api").service(x); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := findCallRoot(root, src); got != "web::scope" {
		t.Fatalf("findCallRoot = %q, want web::scope", got)
	}
}
