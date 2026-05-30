//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsParentWellKnown_NotWellKnown 테스트
package fastapi

import "testing"

func TestIsParentWellKnown_NotWellKnown(t *testing.T) {
	src := []byte("class Base(Other):\n    id: int\n")
	root, _ := parsePython(src)
	if isParentWellKnown(root, src, "Base") {
		t.Fatal("expected false")
	}
}
