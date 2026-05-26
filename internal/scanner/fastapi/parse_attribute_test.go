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
