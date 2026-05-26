//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what isBaseModelSubclass 테스트
package fastapi

import "testing"

func TestIsBaseModelSubclass(t *testing.T) {
	src := []byte("class User(BaseModel):\n    pass\nclass Other:\n    pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_definition")
	if len(classes) < 2 {
		t.Fatalf("expected 2 classes, got %d", len(classes))
	}
	if !isBaseModelSubclass(classes[0], src) {
		t.Fatal("User should be BaseModel subclass")
	}
	if isBaseModelSubclass(classes[1], src) {
		t.Fatal("Other should not be BaseModel subclass")
	}
}
