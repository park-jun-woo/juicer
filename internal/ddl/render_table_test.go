//ff:func feature=ddl type=test control=sequence
//ff:what TestRenderTable_Basic 테스트
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
