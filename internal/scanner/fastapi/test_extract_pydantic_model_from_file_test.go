//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticModel_FromFile 테스트
package fastapi

import "testing"

func TestExtractPydanticModel_FromFile(t *testing.T) {
	dir := t.TempDir()
	p := mkFile(t, dir, "models.py", "class User(BaseModel):\n    id: int\n    name: str\n")
	fields, err := extractPydanticModel(p, "User")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(fields), fields)
	}
}
