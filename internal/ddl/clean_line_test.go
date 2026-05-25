package ddl

import "testing"

func TestCleanLine_WithComment(t *testing.T) {
	got := cleanLine("  id INT NOT NULL -- primary key  ")
	if got != "id INT NOT NULL" {
		t.Fatalf("got %q", got)
	}
}

func TestCleanLine_NoComment(t *testing.T) {
	got := cleanLine("  id INT  ")
	if got != "id INT" {
		t.Fatalf("got %q", got)
	}
}

func TestCleanLine_Empty(t *testing.T) {
	got := cleanLine("")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}
