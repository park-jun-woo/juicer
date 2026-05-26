//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindNodeIndex_Found 테스트
package nestjs

import "testing"

func TestFindNodeIndex_Found(t *testing.T) {
	src := []byte(`const x = 1; const y = 2;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	if root.ChildCount() < 2 {
		t.Skip("not enough children")
	}
	child := root.Child(1)
	idx := findNodeIndex(root, child)
	if idx != 1 {
		t.Fatalf("expected 1, got %d", idx)
	}
}
