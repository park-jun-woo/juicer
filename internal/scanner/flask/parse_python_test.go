//ff:func feature=scan type=test control=sequence topic=flask
//ff:what parsePython 테스트
package flask

import "testing"

func TestParsePython_Success(t *testing.T) {
	root, err := parsePython([]byte("def f():\n    pass\n"))
	if err != nil {
		t.Fatal(err)
	}
	if root == nil || root.Type() != "module" {
		t.Fatalf("expected module root, got %v", root)
	}
}

func TestParsePython_Empty(t *testing.T) {
	root, err := parsePython([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if root == nil {
		t.Fatal("expected non-nil root for empty source")
	}
}
