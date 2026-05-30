//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what isParentWellKnown: 부모가 BaseModel 상속 true / 미상속 / 미발견 false
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

func TestIsParentWellKnown_NotWellKnown(t *testing.T) {
	src := []byte("class Base(Other):\n    id: int\n")
	root, _ := parsePython(src)
	if isParentWellKnown(root, src, "Base") {
		t.Fatal("expected false")
	}
}

func TestIsParentWellKnown_NotFound(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n")
	root, _ := parsePython(src)
	if isParentWellKnown(root, src, "Missing") {
		t.Fatal("expected false")
	}
}
