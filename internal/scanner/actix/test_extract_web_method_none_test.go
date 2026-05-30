//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractWebMethod_None 테스트
package actix

import "testing"

func TestExtractWebMethod_None(t *testing.T) {

	src := []byte(`fn f() { foo.bar; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fe := firstFieldExpr(root)
	if fe == nil {
		t.Fatal("no field_expression found")
	}
	if got := extractWebMethod(fe, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
