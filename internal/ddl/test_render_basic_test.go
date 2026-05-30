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
