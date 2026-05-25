package ddl

import "testing"

func TestRemoveColumn_Found(t *testing.T) {
	cols := []Column{{Name: "id", Raw: "id INT"}, {Name: "name", Raw: "name TEXT"}}
	result := removeColumn(cols, "name")
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
}

func TestRemoveColumn_NotFound(t *testing.T) {
	cols := []Column{{Name: "id", Raw: "id INT"}}
	result := removeColumn(cols, "email")
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
}

func TestRemoveColumn_Empty(t *testing.T) {
	result := removeColumn(nil, "id")
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
