//ff:func feature=scan type=test control=sequence
//ff:what TestFieldsToSchema_Basic 테스트
package scanner

import "testing"

func TestFieldsToSchema_Basic(t *testing.T) {
	fields := []Field{
		{Name: "Name", Type: "string", JSON: "name", Validate: "required"},
		{Name: "Age", Type: "int"},
	}
	schema := fieldsToSchema(fields)
	if schema["type"] != "object" {
		t.Fatal("expected object type")
	}
	if schema["required"] == nil {
		t.Fatal("expected required field")
	}
}

