//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractDecoratorArgs 테스트
package fastapi

import "testing"

func TestExtractDecoratorArgs(t *testing.T) {
	src := []byte("@router.post('/users', status_code=201, response_model=UserOut)\ndef f(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorator")
	}
	callNode := findChildByType(decs[0], "call")
	path, status, respModel := extractDecoratorArgs(callNode, src)
	if path != "/users" {
		t.Fatalf("path: got %q", path)
	}
	if status != 201 {
		t.Fatalf("status: got %d", status)
	}
	if respModel != "UserOut" {
		t.Fatalf("respModel: got %q", respModel)
	}

	// nil callNode
	p, s, r := extractDecoratorArgs(nil, src)
	if p != "" || s != 0 || r != "" {
		t.Fatal("expected empty for nil")
	}
}
