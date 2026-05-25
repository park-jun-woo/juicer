//ff:func feature=scan type=extract control=sequence
//ff:what TestFieldsToSchema_WithRequired 테스트
package scanner

import "testing"

func TestFieldsToSchema_WithRequired(t *testing.T) {
	fields := []Field{
		{Name: "Email", Type: "string", JSON: "email", Validate: "required"},
	}
	schema := fieldsToSchema(fields)
	req, ok := schema["required"].([]string)
	if !ok || len(req) != 1 {
		t.Fatal("expected required field")
	}
}
