//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindParentFieldsInFile_NotFound 테스트
package fastapi

import "testing"

func TestFindParentFieldsInFile_NotFound(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n")
	root, _ := parsePython(src)
	if fields := findParentFieldsInFile(root, src, "Missing", map[string]bool{}); fields != nil {
		t.Fatalf("expected nil, got %+v", fields)
	}
}
