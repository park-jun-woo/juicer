//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractParams_NoParameters 테스트
package fastapi

import "testing"

func TestExtractParams_NoParameters(t *testing.T) {

	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	ri := &routeInfo{path: "/x"}
	extractParams(root, src, ri, nil)
	if len(ri.params) != 0 || len(ri.query) != 0 {
		t.Fatalf("expected no params, got %+v", ri)
	}
}
