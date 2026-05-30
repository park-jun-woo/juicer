//ff:func feature=ddl type=test control=sequence
//ff:what TestRender_Basic 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestRender_Basic(t *testing.T) {
	tables := map[string]*Table{
		"users":  {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
		"orders": {Name: "orders", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	out := Render(nil, tables)
	if !strings.Contains(out, "CREATE TABLE users") {
		t.Fatalf("expected CREATE TABLE users, got %q", out)
	}
	if !strings.Contains(out, "CREATE TABLE orders") {
		t.Fatalf("expected CREATE TABLE orders, got %q", out)
	}
}

func TestRender_WithEnumsAndTables(t *testing.T) {
	enums := []EnumType{{Name: "status", Values: []string{"a", "b"}}}
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	out := Render(enums, tables)
	if !strings.Contains(out, "status") {
		t.Fatalf("expected enum status in output, got %q", out)
	}
	if !strings.Contains(out, "CREATE TABLE users") {
		t.Fatalf("expected CREATE TABLE users, got %q", out)
	}
	// Enum must precede the table, separated by a blank line.
	idxEnum := strings.Index(out, "status")
	idxTable := strings.Index(out, "CREATE TABLE users")
	if idxEnum >= idxTable {
		t.Fatalf("expected enum to precede table, got %q", out)
	}
}

func TestRender_EnumsOnly(t *testing.T) {
	enums := []EnumType{{Name: "status", Values: []string{"a"}}}
	out := Render(enums, map[string]*Table{})
	if !strings.Contains(out, "status") {
		t.Fatalf("expected enum status in output, got %q", out)
	}
	// No tables — no trailing blank line separator should be emitted.
	if strings.Contains(out, "CREATE TABLE") {
		t.Fatalf("expected no tables, got %q", out)
	}
}
