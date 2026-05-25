package ddl

import (
	"strings"
	"testing"
)

func TestRenderTable_Basic(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}, {Name: "name", Raw: "name TEXT"}}}
	var sb strings.Builder
	renderTable(&sb, tbl)
	out := sb.String()
	if !strings.Contains(out, "CREATE TABLE users") {
		t.Fatalf("unexpected output: %q", out)
	}
	if !strings.Contains(out, "id INT") {
		t.Fatalf("missing column in output: %q", out)
	}
}

func TestRenderTable_WithIndex(t *testing.T) {
	tbl := &Table{
		Name:    "users",
		Columns: []Column{{Name: "id", Raw: "id INT"}},
		Indexes: []string{"CREATE INDEX idx_name ON users (name)"},
	}
	var sb strings.Builder
	renderTable(&sb, tbl)
	out := sb.String()
	if !strings.Contains(out, "CREATE INDEX") {
		t.Fatalf("missing index in output: %q", out)
	}
}

func TestRenderTable_WithConstraint(t *testing.T) {
	tbl := &Table{
		Name:        "users",
		Columns:     []Column{{Name: "id", Raw: "id INT"}},
		Constraints: []string{"FOREIGN KEY (id) REFERENCES other(id)"},
	}
	var sb strings.Builder
	renderTable(&sb, tbl)
	out := sb.String()
	if !strings.Contains(out, "FOREIGN KEY") {
		t.Fatalf("missing constraint: %q", out)
	}
}
