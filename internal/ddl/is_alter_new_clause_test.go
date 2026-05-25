package ddl

import "testing"

func TestIsAlterNewClause_AddColumn(t *testing.T) {
	if !isAlterNewClause("ADD COLUMN name TEXT") {
		t.Fatal("expected true")
	}
}

func TestIsAlterNewClause_DropColumn(t *testing.T) {
	if !isAlterNewClause("DROP COLUMN name") {
		t.Fatal("expected true")
	}
}

func TestIsAlterNewClause_AlterColumn(t *testing.T) {
	if !isAlterNewClause("ALTER COLUMN name SET NOT NULL") {
		t.Fatal("expected true")
	}
}

func TestIsAlterNewClause_Rename(t *testing.T) {
	if !isAlterNewClause("RENAME TO new_name") {
		t.Fatal("expected true")
	}
}

func TestIsAlterNewClause_NotClause(t *testing.T) {
	if isAlterNewClause("name TEXT NOT NULL") {
		t.Fatal("expected false")
	}
}
