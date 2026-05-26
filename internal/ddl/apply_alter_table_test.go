//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_AddColumn 테스트
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

	// RENAME TO
	applyAlterTable(tables, "users", "RENAME TO accounts")
	if _, ok := tables["accounts"]; !ok {
		t.Fatal("expected table renamed to accounts")
	}
	if _, ok := tables["users"]; ok {
		t.Fatal("old table name should be deleted")
	}

	// RENAME TO with non-existent table
	applyAlterTable(tables, "nonexistent", "RENAME TO other")

	// Non-column ALTER (e.g. ADD CONSTRAINT)
	applyAlterTable(tables, "accounts", "ADD CONSTRAINT pk PRIMARY KEY (id)")

	// Nil table for column ALTER
	applyAlterTable(tables, "missing", "ADD COLUMN x INT")

	// DROP COLUMN
	applyAlterTable(tables, "accounts", "DROP COLUMN email")
	if len(tables["accounts"].Columns) != 1 {
		t.Fatalf("expected 1 column after DROP, got %d", len(tables["accounts"].Columns))
	}
}
