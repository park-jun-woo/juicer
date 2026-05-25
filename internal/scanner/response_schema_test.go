package scanner

import "testing"

func TestResponseSchema_NamedType(t *testing.T) {
	resp := Response{TypeName: "User", Fields: []Field{{Name: "ID", Type: "int"}}}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["$ref"] == nil {
		t.Fatal("expected $ref")
	}
}

func TestResponseSchema_GinH(t *testing.T) {
	resp := Response{TypeName: "gin.H", Fields: []Field{{Name: "msg", Type: "string"}}}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatalf("expected object, got %v", result["type"])
	}
}

func TestResponseSchema_NoFields(t *testing.T) {
	resp := Response{TypeName: "SomeType"}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatal("expected object")
	}
}

func TestResponseSchema_NoTypeName(t *testing.T) {
	resp := Response{}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatal("expected object")
	}
}

func TestResponseSchema_SliceType(t *testing.T) {
	resp := Response{TypeName: "[]User", Fields: []Field{{Name: "ID", Type: "int"}}}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "array" {
		t.Fatal("expected array")
	}
}
