//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterTableMulti 다중 clause 및 edge case ALTER TABLE 테스트
package ddl

import (
	"testing"
)

func TestApplyAlterTableMulti(t *testing.T) {
	t.Run("ADD COLUMN to nonexistent table", func(t *testing.T) {
		tables := make(map[string]*Table)
		// Should not panic
		applyAlterTable(tables, "nonexistent", "ADD COLUMN email TEXT")
	})

	t.Run("ADD COLUMN duplicate skipped", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Columns: []Column{{Name: "email", Raw: "email TEXT"}}},
		}
		applyAlterTable(tables, "users", "ADD COLUMN email TEXT")
		if len(tables["users"].Columns) != 1 {
			t.Errorf("expected 1 column (no duplicate), got %d", len(tables["users"].Columns))
		}
	})

	t.Run("empty clause from splitAlterClauses", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id BIGINT"}}},
		}
		// Multiple commas or trailing content that results in empty clauses
		applyAlterTable(tables, "users", "ADD COLUMN x TEXT,  ,  DROP COLUMN x")
		// Should not panic, x should be added then dropped
	})

	t.Run("multi-action ADD and DROP", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Columns: []Column{
				{Name: "id", Raw: "id BIGINT"},
				{Name: "old_col", Raw: "old_col TEXT"},
			}},
		}
		applyAlterTable(tables, "users", "ADD COLUMN new_col INT, DROP COLUMN old_col")
		cols := tables["users"].Columns
		if len(cols) != 2 {
			t.Errorf("expected 2 columns, got %d", len(cols))
		}
		// Should have id and new_col
		hasNew := false
		hasOld := false
		for _, c := range cols {
			if c.Name == "new_col" {
				hasNew = true
			}
			if c.Name == "old_col" {
				hasOld = true
			}
		}
		if !hasNew {
			t.Error("expected new_col to be added")
		}
		if hasOld {
			t.Error("expected old_col to be dropped")
		}
	})
}
