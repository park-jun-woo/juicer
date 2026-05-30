//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findPydanticClass 테스트
package fastapi

import "testing"

func TestFindPydanticClass(t *testing.T) {
	src := []byte("class UserCreate(BaseModel):\n    name: str\n    email: str\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	fields := findPydanticClass(root, src, "UserCreate")
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}

	// class not found
	fields2 := findPydanticClass(root, src, "NonExistent")
	if len(fields2) != 0 {
		t.Fatalf("expected 0, got %d", len(fields2))
	}
}

func TestFindPydanticClass_NotBaseModel(t *testing.T) {
	src := []byte("class Plain:\n    id: int\n")
	root, _ := parsePython(src)
	if fields := findPydanticClass(root, src, "Plain"); fields != nil {
		t.Fatalf("expected nil for non-BaseModel, got %+v", fields)
	}
}
