//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindParentFieldsInFile_Found 테스트
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
