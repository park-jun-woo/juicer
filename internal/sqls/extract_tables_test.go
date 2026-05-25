package sqls

import "testing"

func TestExtractTables_Basic(t *testing.T) {
	tables := extractTables([]string{"SELECT * FROM users"})
	if len(tables) == 0 {
		t.Fatal("expected at least one table")
	}
}

func TestExtractTables_Empty(t *testing.T) {
	tables := extractTables(nil)
	if len(tables) != 0 {
		t.Fatalf("expected 0, got %d", len(tables))
	}
}
