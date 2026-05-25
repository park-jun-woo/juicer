//ff:func feature=ddl type=parse control=sequence
//ff:what TestRenderTable 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestRenderTable(t *testing.T) {
	t.Run("columns only", func(t *testing.T) {
		var sb strings.Builder
		tbl := &Table{
			Name: "users",
			Columns: []Column{
				{Name: "id", Raw: "id BIGINT PRIMARY KEY"},
				{Name: "name", Raw: "name TEXT"},
			},
		}
		renderTable(&sb, tbl)
		got := sb.String()
		if !containsStr(got, "CREATE TABLE users") {
			t.Errorf("missing CREATE TABLE, got:\n%s", got)
		}
		if !containsStr(got, "id BIGINT PRIMARY KEY") {
			t.Errorf("missing column id, got:\n%s", got)
		}
	})

	t.Run("columns and constraints", func(t *testing.T) {
		var sb strings.Builder
		tbl := &Table{
			Name: "orders",
			Columns: []Column{
				{Name: "id", Raw: "id BIGINT"},
				{Name: "user_id", Raw: "user_id BIGINT"},
			},
			Constraints: []string{"FOREIGN KEY (user_id) REFERENCES users(id)"},
		}
		renderTable(&sb, tbl)
		got := sb.String()
		if !containsStr(got, "FOREIGN KEY") {
			t.Errorf("missing constraint, got:\n%s", got)
		}
	})

	t.Run("with indexes", func(t *testing.T) {
		var sb strings.Builder
		tbl := &Table{
			Name:    "users",
			Columns: []Column{{Name: "id", Raw: "id BIGINT"}},
			Indexes: []string{"CREATE INDEX idx_users_id ON users (id)"},
		}
		renderTable(&sb, tbl)
		got := sb.String()
		if !containsStr(got, "CREATE INDEX idx_users_id") {
			t.Errorf("missing index, got:\n%s", got)
		}
	})

	t.Run("empty table", func(t *testing.T) {
		var sb strings.Builder
		tbl := &Table{Name: "empty"}
		renderTable(&sb, tbl)
		got := sb.String()
		if !containsStr(got, "CREATE TABLE empty") {
			t.Errorf("missing CREATE TABLE, got:\n%s", got)
		}
	})

	t.Run("multiple constraints", func(t *testing.T) {
		var sb strings.Builder
		tbl := &Table{
			Name: "orders",
			Columns: []Column{
				{Name: "id", Raw: "id BIGINT"},
			},
			Constraints: []string{
				"PRIMARY KEY (id)",
				"CHECK (id > 0)",
			},
		}
		renderTable(&sb, tbl)
		got := sb.String()
		if !containsStr(got, "PRIMARY KEY") {
			t.Errorf("missing PRIMARY KEY, got:\n%s", got)
		}
		if !containsStr(got, "CHECK") {
			t.Errorf("missing CHECK, got:\n%s", got)
		}
	})

	t.Run("column with inline comment in Raw", func(t *testing.T) {
		var sb strings.Builder
		tbl := &Table{
			Name: "t",
			Columns: []Column{
				{Name: "id", Raw: "id BIGINT -- primary key"},
			},
		}
		renderTable(&sb, tbl)
		got := sb.String()
		// cleanLine should strip the inline comment
		if containsStr(got, "-- primary key") {
			t.Errorf("inline comment should be stripped, got:\n%s", got)
		}
	})
}
