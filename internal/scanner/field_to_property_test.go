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


func TestFieldToProperty_Branches(t *testing.T) {
	min, max := 1, 100
	minL, maxL := 2, 50

	// nullable scalar
	if p := fieldToProperty(Field{Type: "string", Nullable: true}); p["nullable"] != true {
		t.Errorf("nullable not set: %v", p)
	}

	// nullable nested struct
	ns := Field{Type: "Addr", Nullable: true, Fields: []Field{{JSON: "c", Type: "string"}}}
	if p := fieldToProperty(ns); p["nullable"] != true {
		t.Errorf("nested nullable not set: %v", p)
	}

	// array of object -> items $ref
	arrObj := Field{Type: "[]User"}
	p := fieldToProperty(arrObj)
	items, _ := p["items"].(map[string]any)
	if items == nil || items["$ref"] == nil {
		t.Errorf("array of object should have $ref items, got %v", p["items"])
	}

	// type:format rule
	tf := fieldToProperty(Field{Type: "string:date-time"})
	if tf["type"] != "string" || tf["format"] != "date-time" {
		t.Errorf("type:format rule failed: %v", tf)
	}

	// enum + constraints
	full := Field{
		Type:      "integer",
		Enum:      []string{"1", "2"},
		Minimum:   &min,
		Maximum:   &max,
		MinLength: &minL,
		MaxLength: &maxL,
	}
	fp := fieldToProperty(full)
	if fp["enum"] == nil {
		t.Errorf("enum not set")
	}
	if fp["minimum"] != 1 || fp["maximum"] != 100 || fp["minLength"] != 2 || fp["maxLength"] != 50 {
		t.Errorf("constraints wrong: %v", fp)
	}
}
