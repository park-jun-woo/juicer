//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterTable 기본 ALTER TABLE 동작 테스트
package ddl

import (
	"testing"
)

func TestApplyAlterTable(t *testing.T) {
	t.Run("ADD COLUMN", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id BIGINT"}}},
		}
		applyAlterTable(tables, "users", "ADD COLUMN email TEXT")
		if len(tables["users"].Columns) != 2 {
			t.Errorf("expected 2 columns, got %d", len(tables["users"].Columns))
		}
	})

	t.Run("DROP COLUMN", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Columns: []Column{
				{Name: "id", Raw: "id BIGINT"},
				{Name: "email", Raw: "email TEXT"},
			}},
		}
		applyAlterTable(tables, "users", "DROP COLUMN email")
		if len(tables["users"].Columns) != 1 {
			t.Errorf("expected 1 column, got %d", len(tables["users"].Columns))
		}
	})

	t.Run("RENAME TO", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users"},
		}
		applyAlterTable(tables, "users", "RENAME TO accounts")
		if _, ok := tables["users"]; ok {
			t.Error("old table name should be removed")
		}
		if _, ok := tables["accounts"]; !ok {
			t.Error("new table name should exist")
		}
	})

	t.Run("RENAME TO nil table", func(t *testing.T) {
		tables := make(map[string]*Table)
		// No panic expected
		applyAlterTable(tables, "nonexistent", "RENAME TO something")
	})

	t.Run("non-column alter skipped", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id BIGINT"}}},
		}
		// ALTER COLUMN should be skipped (not ADD/DROP COLUMN)
		applyAlterTable(tables, "users", "ALTER COLUMN id SET DEFAULT 0")
		if len(tables["users"].Columns) != 1 {
			t.Errorf("expected 1 column unchanged, got %d", len(tables["users"].Columns))
		}
	})
}
