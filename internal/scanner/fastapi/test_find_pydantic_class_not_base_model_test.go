//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindPydanticClass_NotBaseModel 테스트
package fastapi

import "testing"

func TestFindPydanticClass_NotBaseModel(t *testing.T) {
	src := []byte("class Plain:\n    id: int\n")
	root, _ := parsePython(src)
	if fields := findPydanticClass(root, src, "Plain"); fields != nil {
		t.Fatalf("expected nil for non-BaseModel, got %+v", fields)
	}
}
