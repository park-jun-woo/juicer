//ff:func feature=ddl type=test control=sequence
//ff:what TestRender_WithEnumsAndTables 테스트
package ddl

import (
	"strings"
	"testing"
)

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

	idxEnum := strings.Index(out, "status")
	idxTable := strings.Index(out, "CREATE TABLE users")
	if idxEnum >= idxTable {
		t.Fatalf("expected enum to precede table, got %q", out)
	}
}
