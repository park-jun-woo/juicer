//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticModel_NotFound 테스트
package fastapi

import "testing"

func TestExtractPydanticModel_NotFound(t *testing.T) {
	src := []byte(`
from pydantic import BaseModel

class UserCreate(BaseModel):
    name: str
`)
	fields, err := extractPydanticModelFromSource(src, "DoesNotExist")
	if err != nil {
		t.Fatal(err)
	}
	if fields != nil {
		t.Fatalf("expected nil fields, got %v", fields)
	}
}
