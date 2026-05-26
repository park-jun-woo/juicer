//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findAllByType 테스트
package fastapi

import "testing"

func TestFindAllByType(t *testing.T) {
	src := []byte("x = 1\ny = 2\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	ids := findAllByType(root, "identifier")
	if len(ids) < 2 {
		t.Fatalf("expected >= 2 identifiers, got %d", len(ids))
	}
}
