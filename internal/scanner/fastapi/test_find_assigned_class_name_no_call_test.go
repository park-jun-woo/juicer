//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindAssignedClassName_NoCall 테스트
package fastapi

import "testing"

func TestFindAssignedClassName_NoCall(t *testing.T) {
	src := []byte("x = 5\n")
	root, _ := parsePython(src)
	if got := findAssignedClassName(root, "x", src); got != "" {
		t.Fatalf("got %q", got)
	}
}
