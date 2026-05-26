//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what pydanticFieldsToScannerFields 테스트
package fastapi

import "testing"

func TestPydanticFieldsToScannerFields(t *testing.T) {
	fields := []pydanticField{
		{name: "email", typeName: "str"},
		{name: "age", typeName: "int"},
	}
	got := pydanticFieldsToScannerFields(fields)
	if len(got) != 2 {
		t.Fatalf("expected 2, got %d", len(got))
	}
	if got[0].Name != "email" || got[0].Type != "string" {
		t.Fatalf("field 0: %v", got[0])
	}
	if got[1].Name != "age" || got[1].Type != "integer" {
		t.Fatalf("field 1: %v", got[1])
	}

	// empty
	if pydanticFieldsToScannerFields(nil) != nil {
		t.Fatal("expected nil for empty")
	}
}
