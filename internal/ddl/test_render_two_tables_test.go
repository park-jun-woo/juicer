//ff:func feature=ddl type=test control=sequence
//ff:what TestRender_TwoTables 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestRender_TwoTables(t *testing.T) {
	tables := map[string]*Table{
		"users":  {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
		"orders": {Name: "orders", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	out := Render(nil, tables)
	if !strings.Contains(out, "orders") || !strings.Contains(out, "users") {
		t.Fatalf("expected both tables, got %q", out)
	}
}
