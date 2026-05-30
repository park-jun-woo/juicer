//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindAllByType_NoMatch 테스트
package actix

import "testing"

func TestFindAllByType_NoMatch(t *testing.T) {
	src := []byte(`fn f() {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if nodes := findAllByType(root, "struct_item"); len(nodes) != 0 {
		t.Fatalf("expected no matches, got %d", len(nodes))
	}
}
