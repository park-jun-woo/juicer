package scanner

import "testing"

func TestFieldToProperty_String(t *testing.T) {
	f := Field{Name: "Name", Type: "string"}
	prop := fieldToProperty(f)
	if prop["type"] != "string" {
		t.Fatalf("expected string, got %v", prop["type"])
	}
}

func TestFieldToProperty_NestedStruct(t *testing.T) {
	f := Field{Name: "Addr", Type: "Address", Fields: []Field{{Name: "City", Type: "string"}}}
	prop := fieldToProperty(f)
	if prop["type"] != "object" {
		t.Fatalf("expected object, got %v", prop["type"])
	}
}

func TestFieldToProperty_Array(t *testing.T) {
	f := Field{Name: "Tags", Type: "[]string"}
	prop := fieldToProperty(f)
	if prop["type"] != "array" {
		t.Fatalf("expected array, got %v", prop["type"])
	}
}

func TestFieldToProperty_ArrayOfStruct(t *testing.T) {
	f := Field{Name: "Items", Type: "[]Item", Fields: []Field{{Name: "ID", Type: "int"}}}
	prop := fieldToProperty(f)
	if prop["type"] != "array" {
		t.Fatalf("expected array, got %v", prop["type"])
	}
}

func TestFieldToProperty_Pointer(t *testing.T) {
	f := Field{Name: "Age", Type: "*int"}
	prop := fieldToProperty(f)
	if prop["type"] != "integer" {
		t.Fatalf("expected integer, got %v", prop["type"])
	}
}
