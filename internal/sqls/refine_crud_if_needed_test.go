package sqls

import "testing"

func TestRefineCRUDIfNeeded_Select(t *testing.T) {
	got := refineCRUDIfNeeded("SELECT", []string{"SELECT * FROM users"}, nil)
	if got != "SELECT" {
		t.Fatalf("expected SELECT, got %s", got)
	}
}

func TestRefineCRUDIfNeeded_Exec(t *testing.T) {
	got := refineCRUDIfNeeded("EXEC", []string{"INSERT INTO users"}, nil)
	if got != "INSERT" {
		t.Fatalf("expected INSERT, got %s", got)
	}
}
