package ddl

import "testing"

func TestApplyAlterTable_AddColumn(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyAlterTable(tables, "users", "ADD COLUMN email TEXT")
	if len(tables["users"].Columns) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(tables["users"].Columns))
	}
}

func TestApplyAlterTable_DropColumn(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}, {Name: "email", Raw: "email TEXT"}}},
	}
	applyAlterTable(tables, "users", "DROP COLUMN email")
	if len(tables["users"].Columns) != 1 {
		t.Fatalf("expected 1 column, got %d", len(tables["users"].Columns))
	}
}

func TestApplyAlterTable_Rename(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyAlterTable(tables, "users", "RENAME TO accounts")
	if _, ok := tables["accounts"]; !ok {
		t.Fatal("expected table renamed to accounts")
	}
	if _, ok := tables["users"]; ok {
		t.Fatal("expected old name deleted")
	}
}

func TestApplyAlterTable_NoTable(t *testing.T) {
	tables := map[string]*Table{}
	applyAlterTable(tables, "nonexistent", "ADD COLUMN x INT")
}

func TestApplyAlterTable_NonColumnAlter(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyAlterTable(tables, "users", "ALTER COLUMN id SET NOT NULL")
	if len(tables["users"].Columns) != 1 {
		t.Fatal("columns should be unchanged")
	}
}
