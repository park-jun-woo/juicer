//ff:func feature=scan type=convert control=sequence
//ff:what TestResponseSchema 테스트
package scanner

import (
	"testing"
)

func TestResponseSchema(t *testing.T) {
	t.Run("named type", func(t *testing.T) {
		schemas := map[string]any{}
		s := responseSchema(Response{TypeName: "User", Fields: []Field{{Name: "id", Type: "int"}}}, schemas)
		if s["$ref"] == nil {
			t.Error("expected $ref")
		}
	})

	t.Run("gin.H", func(t *testing.T) {
		schemas := map[string]any{}
		s := responseSchema(Response{TypeName: "gin.H", Fields: []Field{{Name: "msg", Type: "string"}}}, schemas)
		if s["$ref"] != nil {
			t.Error("gin.H should not use $ref")
		}
	})

	t.Run("slice type", func(t *testing.T) {
		schemas := map[string]any{}
		s := responseSchema(Response{TypeName: "[]User", Fields: []Field{{Name: "id", Type: "int"}}}, schemas)
		if s["type"] != "array" {
			t.Error("expected array type")
		}
	})

	t.Run("anonymous fields", func(t *testing.T) {
		schemas := map[string]any{}
		s := responseSchema(Response{Fields: []Field{{Name: "id", Type: "int"}}}, schemas)
		if s["type"] != "object" {
			t.Error("expected object type")
		}
	})

	t.Run("type name only", func(t *testing.T) {
		schemas := map[string]any{}
		s := responseSchema(Response{TypeName: "SomeType"}, schemas)
		if s["$ref"] != "#/components/schemas/someType" {
			t.Errorf("expected $ref to someType, got %v", s)
		}
		if schemas["someType"] == nil {
			t.Error("expected placeholder schema registered")
		}
	})

	t.Run("slice type name only", func(t *testing.T) {
		schemas := map[string]any{}
		s := responseSchema(Response{TypeName: "[]SomeType"}, schemas)
		if s["type"] != "array" {
			t.Errorf("expected array type, got %v", s)
		}
		items, ok := s["items"].(map[string]any)
		if !ok || items["$ref"] != "#/components/schemas/someType" {
			t.Errorf("expected items.$ref to someType, got %v", s)
		}
	})

	t.Run("no info", func(t *testing.T) {
		schemas := map[string]any{}
		s := responseSchema(Response{}, schemas)
		if s["type"] != "object" {
			t.Error("expected object type")
		}
	})
}
