//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParsePython_Empty 테스트
package flask

import "testing"

func TestParsePython_Empty(t *testing.T) {
	root, err := parsePython([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if root == nil {
		t.Fatal("expected non-nil root for empty source")
	}
}
