//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDtoFieldsToScannerFields_Basic 테스트
package nestjs

import "testing"

func TestDtoFieldsToScannerFields_Basic(t *testing.T) {
	fields := []dtoField{
		{name: "name", tsType: "string", optional: false, validators: nil},
		{name: "age", tsType: "number", optional: true, validators: []string{"IsInt", "Min(0)"}},
	}
	result := dtoFieldsToScannerFields(fields)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
	if result[0].Name != "name" || result[0].Type != "string" {
		t.Fatalf("unexpected field 0: %+v", result[0])
	}
	if result[1].Validate != "IsInt,Min(0)" {
		t.Fatalf("expected validators, got %q", result[1].Validate)
	}

	// enum field should be overridden to "string"
	enumFields := []dtoField{
		{name: "status", tsType: "TaskStatus", optional: false, validators: []string{"IsEnum"}},
	}
	enumResult := dtoFieldsToScannerFields(enumFields)
	if enumResult[0].Type != "string" {
		t.Fatalf("expected enum type 'string', got %q", enumResult[0].Type)
	}
}
