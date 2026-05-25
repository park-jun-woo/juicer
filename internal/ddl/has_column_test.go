package ddl

import "testing"

func TestHasColumn_Found(t *testing.T) {
	tbl := &Table{Columns: []Column{{Name: "id"}, {Name: "name"}}}
	if !hasColumn(tbl, "id") {
		t.Fatal("expected true")
	}
}

func TestHasColumn_NotFound(t *testing.T) {
	tbl := &Table{Columns: []Column{{Name: "id"}}}
	if hasColumn(tbl, "email") {
		t.Fatal("expected false")
	}
}

func TestHasColumn_Empty(t *testing.T) {
	tbl := &Table{}
	if hasColumn(tbl, "id") {
		t.Fatal("expected false")
	}
}
