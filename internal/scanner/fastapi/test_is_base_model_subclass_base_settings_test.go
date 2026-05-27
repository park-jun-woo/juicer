//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsBaseModelSubclass_BaseSettings 테스트
package fastapi

import "testing"

func TestIsBaseModelSubclass_BaseSettings(t *testing.T) {
	src := []byte("class Settings(BaseSettings):\n    debug: bool = False\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_definition")
	if len(classes) != 1 {
		t.Fatalf("expected 1 class, got %d", len(classes))
	}
	if !isBaseModelSubclass(classes[0], root, src) {
		t.Fatal("Settings(BaseSettings) should be recognized as BaseModel subclass")
	}
}
