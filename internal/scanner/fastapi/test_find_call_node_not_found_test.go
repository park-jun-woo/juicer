//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindCallNode_NotFound 테스트
package fastapi

import "testing"

func TestFindCallNode_NotFound(t *testing.T) {
	src := []byte("x = 5\n")
	root, _ := parsePython(src)
	assigns := findAllByType(root, "assignment")
	if len(assigns) == 0 {
		t.Fatal("no assignment")
	}
	if findCallNode(assigns[0]) != nil {
		t.Fatal("expected nil for no call child")
	}
}
