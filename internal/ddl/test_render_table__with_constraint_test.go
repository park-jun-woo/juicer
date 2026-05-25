//ff:func feature=ddl type=render control=sequence
//ff:what TestRenderTable_WithConstraint 테스트
package ddl

import (
	"strings"
	"testing"
)

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
