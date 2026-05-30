//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindCallRoot_None 테스트
package actix

import "testing"

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
