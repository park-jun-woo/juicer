//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what nodeText 테스트
package fastapi

import "testing"

func TestNodeText(t *testing.T) {
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	text := nodeText(root, src)
	if text == "" {
		t.Fatal("expected non-empty text")
	}
}
