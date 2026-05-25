//ff:func feature=scan type=convert control=sequence
//ff:what TestFieldToProperty 테스트
package scanner

import (
	"testing"
)

func TestFieldToProperty(t *testing.T) {
	t.Run("nested struct", func(t *testing.T) {
		f := Field{Name: "address", Type: "Address", Fields: []Field{{Name: "city", Type: "string"}}}
		prop := fieldToProperty(f)
		if prop["type"] != "object" {
			t.Error("expected object type for nested struct")
		}
	})

	t.Run("nested array struct", func(t *testing.T) {
		f := Field{Name: "items", Type: "[]Item", Fields: []Field{{Name: "id", Type: "int"}}}
		prop := fieldToProperty(f)
		if prop["type"] != "array" {
			t.Error("expected array type")
		}
	})

	t.Run("array type", func(t *testing.T) {
		f := Field{Name: "tags", Type: "[]string"}
		prop := fieldToProperty(f)
		if prop["type"] != "array" {
			t.Error("expected array type")
		}
	})

	t.Run("pointer type", func(t *testing.T) {
		f := Field{Name: "count", Type: "*int"}
		prop := fieldToProperty(f)
		if prop["type"] != "integer" {
			t.Errorf("expected integer, got %v", prop["type"])
		}
	})

	t.Run("with format", func(t *testing.T) {
		f := Field{Name: "count", Type: "int64"}
		prop := fieldToProperty(f)
		if prop["format"] != "int64" {
			t.Errorf("expected format int64, got %v", prop["format"])
		}
	})
}
