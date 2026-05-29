//ff:func feature=ddl type=parse control=sequence
//ff:what TestRender 테스트
package ddl

import (
	"testing"
)

func TestRender(t *testing.T) {
	t.Run("empty tables", func(t *testing.T) {
		got := Render(nil, map[string]*Table{})
		if got != "" {
			t.Errorf("expected empty string, got: %q", got)
		}
	})

	t.Run("single table", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {
				Name:    "users",
				Columns: []Column{{Name: "id", Raw: "id BIGINT PRIMARY KEY"}},
			},
		}
		got := Render(nil, tables)
		if !containsStr(got, "CREATE TABLE users") {
			t.Errorf("expected CREATE TABLE users in output, got:\n%s", got)
		}
	})

	t.Run("multiple tables sorted", func(t *testing.T) {
		tables := map[string]*Table{
			"zebra": {Name: "zebra", Columns: []Column{{Name: "id", Raw: "id INT"}}},
			"alpha": {Name: "alpha", Columns: []Column{{Name: "id", Raw: "id INT"}}},
		}
		got := Render(nil, tables)
		alphaIdx := indexOfStr(got, "alpha")
		zebraIdx := indexOfStr(got, "zebra")
		if alphaIdx >= zebraIdx {
			t.Error("expected alpha before zebra in output")
		}
	})

	t.Run("table with indexes", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {
				Name:    "users",
				Columns: []Column{{Name: "id", Raw: "id BIGINT"}},
				Indexes: []string{"CREATE INDEX idx_users_id ON users (id)"},
			},
		}
		got := Render(nil, tables)
		if !containsStr(got, "CREATE INDEX") {
			t.Errorf("expected CREATE INDEX in output, got:\n%s", got)
		}
	})

	t.Run("table with constraints", func(t *testing.T) {
		tables := map[string]*Table{
			"orders": {
				Name:        "orders",
				Columns:     []Column{{Name: "id", Raw: "id BIGINT"}, {Name: "user_id", Raw: "user_id BIGINT"}},
				Constraints: []string{"FOREIGN KEY (user_id) REFERENCES users(id)"},
			},
		}
		got := Render(nil, tables)
		if !containsStr(got, "FOREIGN KEY") {
			t.Errorf("expected FOREIGN KEY in output, got:\n%s", got)
		}
	})
}
