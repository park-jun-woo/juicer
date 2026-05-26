//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticModel_WithDefault 테스트
package fastapi

import "testing"

func TestExtractPydanticModel_WithDefault(t *testing.T) {
	src := []byte(`
from pydantic import BaseModel

class UserCreate(BaseModel):
    name: str
    role: str = "user"
`)
	fields, err := extractPydanticModelFromSource(src, "UserCreate")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
	if fields[0].Name != "name" {
		t.Fatalf("expected name, got %s", fields[0].Name)
	}
	if fields[1].Name != "role" {
		t.Fatalf("expected role, got %s", fields[1].Name)
	}
}
