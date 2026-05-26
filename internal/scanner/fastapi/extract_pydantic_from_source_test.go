//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractPydanticModelFromSource 테스트
package fastapi

import "testing"

func TestExtractPydanticModelFromSource(t *testing.T) {
	src := []byte("class User(BaseModel):\n    name: str\n    email: str\n")
	fields, err := extractPydanticModelFromSource(src, "User")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
}
