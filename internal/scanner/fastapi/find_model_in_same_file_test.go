//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findModelInSameFile 테스트
package fastapi

import "testing"

func TestFindModelInSameFile(t *testing.T) {
	src := []byte("class UserCreate(BaseModel):\n    name: str\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	if !findModelInSameFile(root, src, "UserCreate") {
		t.Fatal("expected to find UserCreate")
	}
	if findModelInSameFile(root, src, "NonExistent") {
		t.Fatal("should not find NonExistent")
	}
}

func TestFindModelInSameFile_MultipleClasses(t *testing.T) {
	src := []byte("class A(BaseModel):\n    x: int\nclass B(BaseModel):\n    y: int\n")
	root, _ := parsePython(src)
	if !findModelInSameFile(root, src, "B") {
		t.Fatal("expected B (skips A first)")
	}
}
