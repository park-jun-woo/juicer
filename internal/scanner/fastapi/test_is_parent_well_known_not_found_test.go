//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsParentWellKnown_NotFound 테스트
package fastapi

import "testing"

func TestIsParentWellKnown_NotFound(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n")
	root, _ := parsePython(src)
	if isParentWellKnown(root, src, "Missing") {
		t.Fatal("expected false")
	}
}
