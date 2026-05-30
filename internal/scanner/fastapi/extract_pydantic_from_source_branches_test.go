//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractPydanticModelFromSource: 모델 발견 / 미발견 (parse err은 도달불가)
package fastapi

import "testing"

func TestExtractPydanticModelFromSource_Found(t *testing.T) {
	src := []byte("class User(BaseModel):\n    id: int\n")
	fields, err := extractPydanticModelFromSource(src, "User")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 || fields[0].Name != "id" {
		t.Fatalf("got %+v", fields)
	}
}

func TestExtractPydanticModelFromSource_NotFound(t *testing.T) {
	src := []byte("class Other(BaseModel):\n    x: int\n")
	fields, err := extractPydanticModelFromSource(src, "Missing")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 0 {
		t.Fatalf("expected no fields for missing class, got %+v", fields)
	}
	// parse-error branch is unreachable: tree-sitter parses any byte slice.
}
