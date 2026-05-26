//ff:func feature=scan type=test control=sequence
//ff:what TestBodySchema_Named 테스트
package scanner

import "testing"

func TestBodySchema_Named(t *testing.T) {
	body := &Body{TypeName: "User", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}

	// slice type
	sliceBody := &Body{TypeName: "[]Item", Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}}
	result = bodySchema(sliceBody, schemas)
	if result["type"] != "array" {
		t.Fatalf("expected array type, got %v", result)
	}

	// anonymous fields (no TypeName)
	anonBody := &Body{Fields: []Field{{Name: "x", JSON: "x", Type: "string"}}}
	result = bodySchema(anonBody, schemas)
	if result == nil {
		t.Fatal("expected non-nil for anonymous")
	}

	// empty body
	emptyBody := &Body{}
	result = bodySchema(emptyBody, schemas)
	if result["type"] != "object" {
		t.Fatalf("expected object type, got %v", result)
	}
}
