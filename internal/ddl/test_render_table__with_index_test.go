//ff:func feature=ddl type=render control=sequence
//ff:what TestRenderTable_WithIndex 테스트
package ddl

import (
	"strings"
	"testing"
)

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
