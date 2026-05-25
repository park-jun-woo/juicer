//ff:func feature=scan type=convert control=sequence
//ff:what TestBodySchema 테스트
package scanner

import (
	"testing"
)

func TestBodySchema(t *testing.T) {
	t.Run("named type with fields", func(t *testing.T) {
		body := &Body{
			TypeName: "User",
			Fields:   []Field{{Name: "name", Type: "string"}},
		}
		schemas := map[string]any{}
		s := bodySchema(body, schemas)
		if s["$ref"] == nil {
			t.Error("expected $ref")
		}
	})

	t.Run("slice named type", func(t *testing.T) {
		body := &Body{
			TypeName: "[]User",
			Fields:   []Field{{Name: "name", Type: "string"}},
		}
		schemas := map[string]any{}
		s := bodySchema(body, schemas)
		if s["type"] != "array" {
			t.Error("expected array type")
		}
	})

	t.Run("anonymous fields", func(t *testing.T) {
		body := &Body{
			Fields: []Field{{Name: "name", Type: "string"}},
		}
		schemas := map[string]any{}
		s := bodySchema(body, schemas)
		if s["type"] != "object" {
			t.Error("expected object type")
		}
	})

	t.Run("no fields", func(t *testing.T) {
		body := &Body{}
		schemas := map[string]any{}
		s := bodySchema(body, schemas)
		if s["type"] != "object" {
			t.Error("expected object type")
		}
	})
}
