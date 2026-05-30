//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticFields 테스트
package fastapi

import "testing"

func TestExtractPydanticFields(t *testing.T) {
	src := []byte("class UserCreate(BaseModel):\n    name: str\n    email: str\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_definition")
	if len(classes) == 0 {
		t.Fatal("no class_definition")
	}
	fields := extractPydanticFields(classes[0], src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
	if fields[0].name != "name" || fields[1].name != "email" {
		t.Fatalf("unexpected fields: %+v", fields)
	}
}
