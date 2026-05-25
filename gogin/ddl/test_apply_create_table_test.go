//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyCreateTable 기본 CREATE TABLE 동작 테스트
package ddl

import (
	"testing"
)

func TestApplyCreateTable(t *testing.T) {
	t.Run("basic table", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE users (id BIGINT PRIMARY KEY, name TEXT NOT NULL)"
		applyCreateTable(tables, "users", stmt)

		tbl, ok := tables["users"]
		if !ok {
			t.Fatal("table 'users' not found")
		}
		if len(tbl.Columns) != 2 {
			t.Errorf("expected 2 columns, got %d", len(tbl.Columns))
		}
	})

	t.Run("table with constraints", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE orders (id BIGINT, user_id BIGINT, FOREIGN KEY (user_id) REFERENCES users(id))"
		applyCreateTable(tables, "orders", stmt)

		tbl := tables["orders"]
		if len(tbl.Columns) != 2 {
			t.Errorf("expected 2 columns, got %d", len(tbl.Columns))
		}
		if len(tbl.Constraints) != 1 {
			t.Errorf("expected 1 constraint, got %d", len(tbl.Constraints))
		}
	})

	t.Run("empty body", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE empty"
		applyCreateTable(tables, "empty", stmt)

		tbl, ok := tables["empty"]
		if !ok {
			t.Fatal("table 'empty' not found")
		}
		if len(tbl.Columns) != 0 {
			t.Errorf("expected 0 columns, got %d", len(tbl.Columns))
		}
	})

	t.Run("case insensitive name", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE Users (id INT)"
		applyCreateTable(tables, "Users", stmt)

		if _, ok := tables["users"]; !ok {
			t.Error("table name should be lowercased")
		}
	})

	t.Run("table with inline comments", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE users (\n    id BIGINT, -- primary key\n    name TEXT -- user name\n)"
		applyCreateTable(tables, "users", stmt)

		tbl := tables["users"]
		if len(tbl.Columns) != 2 {
			t.Errorf("expected 2 columns, got %d", len(tbl.Columns))
		}
	})
}
