//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what parseAttribute 테스트
package fastapi

import "testing"

func TestParseAttribute(t *testing.T) {
	src := []byte("router.get('/x')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	attrs := findAllByType(root, "attribute")
	if len(attrs) == 0 {
		t.Fatal("no attribute nodes")
	}
	varName, method := parseAttribute(attrs[0], src)
	if varName != "router" || method != "GET" {
		t.Fatalf("got var=%q method=%q", varName, method)
	}
}

func TestParseAttribute_NotHTTP(t *testing.T) {
	src := []byte("router.foo('/x')\n")
	root, _ := parsePython(src)
	attrs := findAllByType(root, "attribute")
	v, m := parseAttribute(attrs[0], src)
	if v != "" || m != "" {
		t.Fatalf("expected empty, got v=%q m=%q", v, m)
	}
}

func TestParseAttribute_NoDot(t *testing.T) {
	// SplitN with no dot -> len(parts) != 2. Pass a non-attribute node
	// (an identifier) whose text has no dot.
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
