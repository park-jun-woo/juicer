//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what convertPropertyPair 테스트
package fastify

import "testing"

func TestConvertPropertyPair_Required(t *testing.T) {
	pairs, src := schemaPairs(t, `{ name: { type: "string" }, age: { type: "integer" } }`)
	requiredSet := map[string]bool{"name": true}

	var nameField, ageField bool
	for _, p := range pairs {
		f := convertPropertyPair(p, src, requiredSet)
		if f == nil {
			t.Fatal("nil field")
		}
		switch f.Name {
		case "name":
			nameField = true
			if f.Validate != "required" {
				t.Errorf("name should be required, got %q", f.Validate)
			}
		case "age":
			ageField = true
			if f.Validate == "required" {
				t.Error("age should not be required")
			}
		}
	}
	if !nameField || !ageField {
		t.Errorf("missing fields name=%v age=%v", nameField, ageField)
	}
}

func TestConvertPropertyPair_NotRequiredEmptySet(t *testing.T) {
	pairs, src := schemaPairs(t, `{ title: { type: "string" } }`)
	f := convertPropertyPair(pairs[0], src, map[string]bool{})
	if f == nil || f.Name != "title" || f.Validate == "required" {
		t.Fatalf("unexpected field: %+v", f)
	}
}
