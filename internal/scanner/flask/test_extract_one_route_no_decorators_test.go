//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractOneRoute_NoDecorators 테스트
package flask

import "testing"

func TestExtractOneRoute_NoDecorators(t *testing.T) {

	b := []byte("def plain():\n    pass\n")
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	fns := findAllByType(root, "function_definition")
	if len(fns) == 0 {
		t.Fatal("no function_definition")
	}
	if routes := extractOneRoute(fns[0], b, make(blueprintPrefix), "app.py"); routes != nil {
		t.Fatalf("non-decorated def should yield nil, got %v", routes)
	}
}
