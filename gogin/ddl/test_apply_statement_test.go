//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyStatement 테스트
package ddl

import (
	"testing"
)

func TestApplyStatement(t *testing.T) {
	t.Run("CREATE TABLE", func(t *testing.T) {
		tables := make(map[string]*Table)
		applyStatement(tables, "CREATE TABLE users (id BIGINT PRIMARY KEY, name TEXT)")
		if _, ok := tables["users"]; !ok {
			t.Error("expected table 'users' to be created")
		}
	})

	t.Run("CREATE TABLE with leading comments", func(t *testing.T) {
		tables := make(map[string]*Table)
		applyStatement(tables, "-- migration\nCREATE TABLE items (id INT)")
		if _, ok := tables["items"]; !ok {
			t.Error("expected table 'items' to be created")
		}
	})

	t.Run("DROP TABLE", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users"},
		}
		applyStatement(tables, "DROP TABLE users")
		if _, ok := tables["users"]; ok {
			t.Error("expected table 'users' to be dropped")
		}
	})

	t.Run("DROP TABLE IF EXISTS", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users"},
		}
		applyStatement(tables, "DROP TABLE IF EXISTS users")
		if _, ok := tables["users"]; ok {
			t.Error("expected table 'users' to be dropped")
		}
	})

	t.Run("ALTER TABLE", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id BIGINT"}}},
		}
		applyStatement(tables, "ALTER TABLE users ADD COLUMN email TEXT")
		tbl := tables["users"]
		if len(tbl.Columns) != 2 {
			t.Errorf("expected 2 columns, got %d", len(tbl.Columns))
		}
	})

	t.Run("CREATE INDEX", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users"},
		}
		applyStatement(tables, "CREATE INDEX idx_users_name ON users (name)")
		if len(tables["users"].Indexes) != 1 {
			t.Errorf("expected 1 index, got %d", len(tables["users"].Indexes))
		}
	})

	t.Run("DROP INDEX", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users", Indexes: []string{"CREATE INDEX idx_users_name ON users (name)"}},
		}
		applyStatement(tables, "DROP INDEX idx_users_name")
		if len(tables["users"].Indexes) != 0 {
			t.Errorf("expected 0 indexes, got %d", len(tables["users"].Indexes))
		}
	})

	t.Run("unrecognized statement", func(t *testing.T) {
		tables := make(map[string]*Table)
		// Should not panic on unrecognized statements
		applyStatement(tables, "INSERT INTO users VALUES (1, 'test')")
		if len(tables) != 0 {
			t.Errorf("expected no tables, got %d", len(tables))
		}
	})
}
