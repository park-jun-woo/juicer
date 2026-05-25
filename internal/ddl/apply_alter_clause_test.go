package ddl

import "testing"

func TestApplyAlterClause_AddColumn(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}}
	applyAlterClause(tbl, "ADD COLUMN name TEXT NOT NULL")
	if len(tbl.Columns) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(tbl.Columns))
	}
}

func TestApplyAlterClause_DropColumn(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}, {Name: "name", Raw: "name TEXT"}}}
	applyAlterClause(tbl, "DROP COLUMN name")
	if len(tbl.Columns) != 1 {
		t.Fatalf("expected 1 column, got %d", len(tbl.Columns))
	}
}

func TestApplyAlterClause_Empty(t *testing.T) {
	tbl := &Table{Name: "users"}
	applyAlterClause(tbl, "")
}

func TestApplyAlterClause_DuplicateColumn(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "name", Raw: "name TEXT"}}}
	applyAlterClause(tbl, "ADD COLUMN name TEXT")
	if len(tbl.Columns) != 1 {
		t.Fatalf("expected 1 column (no duplicate), got %d", len(tbl.Columns))
	}
}
