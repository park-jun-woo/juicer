//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsBaseModelSubclass_SameFileInheritance 테스트
package fastapi

import "testing"

func TestIsBaseModelSubclass_SameFileInheritance(t *testing.T) {
	src := []byte("class UserBase(BaseModel):\n    email: str\nclass UserCreate(UserBase):\n    password: str\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_definition")
	if len(classes) != 2 {
		t.Fatalf("expected 2 classes, got %d", len(classes))
	}
	if !isBaseModelSubclass(classes[0], root, src) {
		t.Fatal("UserBase(BaseModel) should be recognized")
	}
	if !isBaseModelSubclass(classes[1], root, src) {
		t.Fatal("UserCreate(UserBase) should be recognized via same-file inheritance")
	}
}
