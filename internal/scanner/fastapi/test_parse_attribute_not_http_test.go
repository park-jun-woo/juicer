//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseAttribute_NotHTTP 테스트
package fastapi

import "testing"

func TestParseAttribute_NotHTTP(t *testing.T) {
	src := []byte("router.foo('/x')\n")
	root, _ := parsePython(src)
	attrs := findAllByType(root, "attribute")
	v, m := parseAttribute(attrs[0], src)
	if v != "" || m != "" {
		t.Fatalf("expected empty, got v=%q m=%q", v, m)
	}
}
