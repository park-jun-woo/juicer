package ddl

import "testing"

func TestApplyCreateTable_Basic(t *testing.T) {
	tables := make(map[string]*Table)
	applyCreateTable(tables, "users", "CREATE TABLE users (id INT PRIMARY KEY, name TEXT NOT NULL)")
	tbl := tables["users"]
	if tbl == nil {
		t.Fatal("expected table")
	}
	if len(tbl.Columns) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(tbl.Columns))
	}
}

func TestApplyCreateTable_WithConstraints(t *testing.T) {
	tables := make(map[string]*Table)
	stmt := "CREATE TABLE users (id INT, name TEXT, FOREIGN KEY (id) REFERENCES other(id))"
	applyCreateTable(tables, "users", stmt)
	tbl := tables["users"]
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
}

func TestApplyCreateTable_NoParen(t *testing.T) {
	tables := make(map[string]*Table)
	applyCreateTable(tables, "empty", "CREATE TABLE empty")
	if tables["empty"] == nil {
		t.Fatal("expected table even without parens")
	}
}
