//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_String 테스트
package scanner

import "testing"

func TestFieldToProperty_String(t *testing.T) {
	f := Field{Name: "Name", Type: "string"}
	prop := fieldToProperty(f)
	if prop["type"] != "string" {
		t.Fatalf("expected string, got %v", prop["type"])
	}

	// nested struct
	nested := Field{Name: "Address", Type: "Address", Fields: []Field{{Name: "city", JSON: "city", Type: "string"}}}
	prop = fieldToProperty(nested)
	if prop["type"] != "object" {
		t.Fatalf("expected object for nested struct, got %v", prop["type"])
	}

	// nested struct slice
	nestedSlice := Field{Name: "Items", Type: "[]Item", Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}}
	prop = fieldToProperty(nestedSlice)
	if prop["type"] != "array" {
		t.Fatalf("expected array for nested slice, got %v", prop["type"])
	}

	// array type
	arr := Field{Name: "Tags", Type: "[]string"}
	prop = fieldToProperty(arr)
	if prop["type"] != "array" {
		t.Fatalf("expected array, got %v", prop["type"])
	}

	// pointer type
	ptr := Field{Name: "Age", Type: "*int"}
	prop = fieldToProperty(ptr)
	if prop["type"] != "integer" {
		t.Fatalf("expected integer for *int, got %v", prop["type"])
	}

	// format hint (time.Time -> date-time)
	ts := Field{Name: "CreatedAt", Type: "time.Time"}
	prop = fieldToProperty(ts)
	if prop["format"] != "date-time" {
		t.Fatalf("expected date-time format, got %v", prop["format"])
	}
}

