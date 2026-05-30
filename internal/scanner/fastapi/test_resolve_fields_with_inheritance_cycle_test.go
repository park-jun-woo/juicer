//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveFieldsWithInheritance_Cycle 테스트
package fastapi

import "testing"

func TestResolveFieldsWithInheritance_Cycle(t *testing.T) {
	src := []byte("class A(B):\n    x: int\n")
	root, _ := parsePython(src)
	a := classByName(root, src, "A")

	if got := resolveFieldsWithInheritance(a, root, src, map[string]bool{"A": true}); got != nil {
		t.Fatalf("expected nil for visited cycle, got %v", got)
	}
}
