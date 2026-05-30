//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractWebMethod 테스트
package actix

import "testing"

func TestExtractWebMethod(t *testing.T) {

	src := []byte(`fn f() { web::get().to(h); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fe := firstFieldExpr(root)
	if fe == nil {
		t.Fatal("no field_expression found")
	}
	if got := extractWebMethod(fe, src); got != "GET" {
		t.Fatalf("extractWebMethod = %q, want GET", got)
	}
}
