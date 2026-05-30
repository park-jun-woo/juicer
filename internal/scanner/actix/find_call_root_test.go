//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findCallRoot — web::scope/web::resource 루트 식별을 검증
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

func TestFindCallRoot_None(t *testing.T) {
	src := []byte(`fn f() { other::thing("/x"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := findCallRoot(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
