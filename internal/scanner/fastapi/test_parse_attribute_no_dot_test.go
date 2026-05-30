//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseAttribute_NoDot 테스트
package fastapi

import "testing"

func TestParseAttribute_NoDot(t *testing.T) {

	src := []byte("router\n")
	root, _ := parsePython(src)
	ids := findAllByType(root, "identifier")
	if len(ids) == 0 {
		t.Fatal("no identifier")
	}
	v, m := parseAttribute(ids[0], src)
	if v != "" || m != "" {
		t.Fatalf("expected empty, got v=%q m=%q", v, m)
	}
}
