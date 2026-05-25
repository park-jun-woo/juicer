//ff:func feature=ddl type=render control=sequence
//ff:what TestRender_MultipleTables 테스트
package ddl

import (
	"strings"
	"testing"
)

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
