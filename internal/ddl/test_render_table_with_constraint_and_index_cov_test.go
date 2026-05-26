//ff:func feature=ddl type=test control=sequence
//ff:what TestRenderTable_WithConstraintAndIndexCov 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestRenderTable_WithConstraintAndIndexCov(t *testing.T) {
	tbl := &Table{
		Name:        "orders",
		Columns:     []Column{{Name: "id", Raw: "id INT"}, {Name: "total", Raw: "total DECIMAL"}},
		Constraints: []string{"CONSTRAINT pk PRIMARY KEY (id)"},
		Indexes:     []string{"CREATE INDEX idx_total ON orders (total)"},
	}
	var sb strings.Builder
	renderTable(&sb, tbl)
	out := sb.String()
	if !strings.Contains(out, "CONSTRAINT pk") {
		t.Fatalf("missing constraint in output: %q", out)
	}
	if !strings.Contains(out, "CREATE INDEX idx_total") {
		t.Fatalf("missing index in output: %q", out)
	}
}
