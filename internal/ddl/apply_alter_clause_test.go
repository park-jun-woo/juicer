//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterClause_AddColumn 테스트
package ddl

import "testing"

func TestApplyAlterClause_AddColumn(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}}
	applyAlterClause(tbl, "ADD COLUMN name TEXT NOT NULL")
	if len(tbl.Columns) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(tbl.Columns))
	}
}
