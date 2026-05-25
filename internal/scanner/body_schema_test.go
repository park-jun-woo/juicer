package scanner

import "testing"

func TestBodySchema_Named(t *testing.T) {
	body := &Body{TypeName: "User", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBodySchema_Slice(t *testing.T) {
	body := &Body{TypeName: "[]User", Fields: []Field{{Name: "id", JSON: "id"}}}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBodySchema_NoFields(t *testing.T) {
	body := &Body{}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
