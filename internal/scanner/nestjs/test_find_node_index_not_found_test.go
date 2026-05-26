//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindNodeIndex_NotFound 테스트
package nestjs

import "testing"

func TestFindNodeIndex_NotFound(t *testing.T) {
	src1 := []byte(`const x = 1;`)
	src2 := []byte(`const y = 2;`)
	root1, _ := parseTypeScript(src1)
	root2, _ := parseTypeScript(src2)
	child := root2.Child(0)
	idx := findNodeIndex(root1, child)
	if idx != -1 {
		t.Fatalf("expected -1, got %d", idx)
	}
}
