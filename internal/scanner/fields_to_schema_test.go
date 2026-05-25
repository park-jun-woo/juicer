package scanner

import "testing"

func TestFieldsToSchema_Basic(t *testing.T) {
	fields := []Field{
		{Name: "Name", Type: "string", JSON: "name"},
		{Name: "Age", Type: "int", JSON: "age"},
	}
	schema := fieldsToSchema(fields)
	if schema["type"] != "object" {
		t.Fatal("expected object type")
	}
}

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

func TestFieldsToSchema_NoJSON(t *testing.T) {
	fields := []Field{{Name: "ID", Type: "int"}}
	schema := fieldsToSchema(fields)
	props := schema["properties"].(map[string]any)
	if _, ok := props["ID"]; !ok {
		t.Fatal("expected ID in properties")
	}
}
