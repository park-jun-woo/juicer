package ddl

import (
	"strings"
	"testing"
)

func TestRender_Basic(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	out := Render(tables)
	if !strings.Contains(out, "CREATE TABLE users") {
		t.Fatalf("expected CREATE TABLE, got %q", out)
	}
}

func TestRender_Empty(t *testing.T) {
	out := Render(map[string]*Table{})
	if out != "" {
		t.Fatalf("expected empty, got %q", out)
	}
}

func TestRender_MultipleTables(t *testing.T) {
	tables := map[string]*Table{
		"a": {Name: "a", Columns: []Column{{Name: "id", Raw: "id INT"}}},
		"b": {Name: "b", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	out := Render(tables)
	if !strings.Contains(out, "CREATE TABLE a") || !strings.Contains(out, "CREATE TABLE b") {
		t.Fatalf("expected both tables, got %q", out)
	}
}
