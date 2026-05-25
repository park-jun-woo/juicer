package sqls

import "testing"

func TestRefineCRUD_Insert(t *testing.T) {
	got := refineCRUD([]string{"INSERT INTO users"})
	if got != "INSERT" {
		t.Fatalf("expected INSERT, got %s", got)
	}
}

func TestRefineCRUD_Update(t *testing.T) {
	got := refineCRUD([]string{"UPDATE users SET name = $1"})
	if got != "UPDATE" {
		t.Fatalf("expected UPDATE, got %s", got)
	}
}

func TestRefineCRUD_Delete(t *testing.T) {
	got := refineCRUD([]string{"DELETE FROM users"})
	if got != "DELETE" {
		t.Fatalf("expected DELETE, got %s", got)
	}
}

func TestRefineCRUD_Default(t *testing.T) {
	got := refineCRUD([]string{"CALL proc()"})
	if got != "EXEC" {
		t.Fatalf("expected EXEC, got %s", got)
	}
}
