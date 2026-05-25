package sqls

import "testing"

func TestExtractTablesFromSQL_Select(t *testing.T) {
	tables := extractTablesFromSQL("SELECT * FROM users WHERE id = $1")
	found := false
	for _, tbl := range tables {
		if tbl == "users" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected users table, got %v", tables)
	}
}

func TestExtractTablesFromSQL_Insert(t *testing.T) {
	tables := extractTablesFromSQL("INSERT INTO orders (id) VALUES ($1)")
	found := false
	for _, tbl := range tables {
		if tbl == "orders" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected orders table, got %v", tables)
	}
}
