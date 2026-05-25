package ddl

import "testing"

func TestExtractColumnName_Basic(t *testing.T) {
	got := extractColumnName("id INT PRIMARY KEY")
	if got != "id" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractColumnName_Empty(t *testing.T) {
	got := extractColumnName("")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractColumnName_WithComment(t *testing.T) {
	got := extractColumnName("-- comment\nname TEXT")
	if got != "name" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractColumnName_OnlyComment(t *testing.T) {
	got := extractColumnName("-- just a comment")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}
