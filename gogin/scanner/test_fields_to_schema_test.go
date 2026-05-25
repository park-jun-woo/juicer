//ff:func feature=scan type=convert control=sequence
//ff:what TestFieldsToSchema 테스트
package scanner

import (
	"testing"
)

func TestFieldsToSchema(t *testing.T) {
	fields := []Field{
		{Name: "id", JSON: "id", Type: "int", Validate: "required"},
		{Name: "name", JSON: "name", Type: "string"},
		{Name: "Email", Type: "string"}, // no JSON tag, use Name
	}

	schema := fieldsToSchema(fields)
	if schema["type"] != "object" {
		t.Error("expected object type")
	}
	props := schema["properties"].(map[string]any)
	if len(props) != 3 {
		t.Errorf("expected 3 properties, got %d", len(props))
	}
	required := schema["required"].([]string)
	if len(required) != 1 || required[0] != "id" {
		t.Errorf("expected required=[id], got %v", required)
	}
}
