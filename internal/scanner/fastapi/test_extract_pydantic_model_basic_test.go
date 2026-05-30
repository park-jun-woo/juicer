//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticModel_Basic 테스트
package fastapi

import "testing"

func TestExtractPydanticModel_Basic(t *testing.T) {
	src := []byte(`
from pydantic import BaseModel

class UserCreate(BaseModel):
    name: str
    email: str
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
	if fields[0].Type != "string" {
		t.Fatalf("expected string, got %s", fields[0].Type)
	}
	if fields[1].Name != "email" {
		t.Fatalf("expected email, got %s", fields[1].Name)
	}
}
