//ff:func feature=ddl type=test control=sequence
//ff:what TestRenderTable_ColumnsOnly_LastNoComma 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestRenderTable_ColumnsOnly_LastNoComma(t *testing.T) {

	tbl := &Table{
		Name:    "users",
		Columns: []Column{{Name: "id", Raw: "id INT"}, {Name: "name", Raw: "name TEXT"}},
	}
	var sb strings.Builder
	renderTable(&sb, tbl)
	out := sb.String()
	if !strings.Contains(out, "name TEXT\n)") {
		t.Fatalf("expected last column without trailing comma, got %q", out)
	}
	if strings.Contains(out, "name TEXT,") {
		t.Fatalf("last column should not have trailing comma: %q", out)
	}
}
