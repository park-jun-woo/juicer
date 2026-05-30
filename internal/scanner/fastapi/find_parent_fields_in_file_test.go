//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findParentFieldsInFile: 클래스 발견 필드반환 / 미발견 nil
package fastapi

import "testing"

func TestFindParentFieldsInFile_Found(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n    name: str\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	fields := findParentFieldsInFile(root, src, "Base", map[string]bool{})
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(fields), fields)
	}
}

func TestFindParentFieldsInFile_NotFound(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n")
	root, _ := parsePython(src)
	if fields := findParentFieldsInFile(root, src, "Missing", map[string]bool{}); fields != nil {
		t.Fatalf("expected nil, got %+v", fields)
	}
}
