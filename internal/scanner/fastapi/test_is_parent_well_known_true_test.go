//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsParentWellKnown_True 테스트
package fastapi

import "testing"

func TestIsParentWellKnown_True(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\nclass Child(Base):\n    name: str\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	if !isParentWellKnown(root, src, "Base") {
		t.Fatal("expected true")
	}
}
