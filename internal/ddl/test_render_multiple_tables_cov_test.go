//ff:func feature=ddl type=test control=sequence
//ff:what TestRender_MultipleTablesCov 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestRender_MultipleTablesCov(t *testing.T) {
	tables := map[string]*Table{
		"users":  {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
		"orders": {Name: "orders", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	out := Render(tables)
	if !strings.Contains(out, "CREATE TABLE orders") || !strings.Contains(out, "CREATE TABLE users") {
		t.Fatalf("expected both tables, got %q", out)
	}
}
