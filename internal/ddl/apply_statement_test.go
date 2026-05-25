package ddl

import "testing"

func TestApplyStatement_CreateTable(t *testing.T) {
	tables := make(map[string]*Table)
	applyStatement(tables, "CREATE TABLE users (id INT PRIMARY KEY)")
	if tables["users"] == nil {
		t.Fatal("expected users table")
	}
}

func TestApplyStatement_CreateIndex(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyStatement(tables, "CREATE INDEX idx_name ON users (name)")
	if len(tables["users"].Indexes) != 1 {
		t.Fatal("expected 1 index")
	}
}

func TestApplyStatement_DropIndex(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Indexes: []string{"CREATE INDEX idx_name ON users (name)"}},
	}
	applyStatement(tables, "DROP INDEX idx_name")
	if len(tables["users"].Indexes) != 0 {
		t.Fatal("expected 0 indexes")
	}
}

func TestApplyStatement_DropTable(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyStatement(tables, "DROP TABLE users")
	if _, ok := tables["users"]; ok {
		t.Fatal("expected users to be deleted")
	}
}

func TestApplyStatement_AlterTable(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyStatement(tables, "ALTER TABLE users ADD COLUMN email TEXT")
	if len(tables["users"].Columns) != 2 {
		t.Fatal("expected 2 columns")
	}
}

func TestApplyStatement_LeadingComment(t *testing.T) {
	tables := make(map[string]*Table)
	applyStatement(tables, "-- comment\nCREATE TABLE items (id INT)")
	if tables["items"] == nil {
		t.Fatal("expected items table")
	}
}
