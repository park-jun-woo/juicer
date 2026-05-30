//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_Branches 테스트
package scanner

import "testing"

func TestFieldToProperty_Branches(t *testing.T) {
	min, max := 1, 100
	minL, maxL := 2, 50

	if p := fieldToProperty(Field{Type: "string", Nullable: true}); p["nullable"] != true {
		t.Errorf("nullable not set: %v", p)
	}

	ns := Field{Type: "Addr", Nullable: true, Fields: []Field{{JSON: "c", Type: "string"}}}
	if p := fieldToProperty(ns); p["nullable"] != true {
		t.Errorf("nested nullable not set: %v", p)
	}

	arrObj := Field{Type: "[]User"}
	p := fieldToProperty(arrObj)
	items, _ := p["items"].(map[string]any)
	if items == nil || items["$ref"] == nil {
		t.Errorf("array of object should have $ref items, got %v", p["items"])
	}

	tf := fieldToProperty(Field{Type: "string:date-time"})
	if tf["type"] != "string" || tf["format"] != "date-time" {
		t.Errorf("type:format rule failed: %v", tf)
	}

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
