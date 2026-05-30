//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindCallRoot_Resource 테스트
package actix

import "testing"

func TestFindCallRoot_Resource(t *testing.T) {
	src := []byte(`fn f() { web::resource("/x"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := findCallRoot(root, src); got != "web::resource" {
		t.Fatalf("findCallRoot = %q, want web::resource", got)
	}
}
