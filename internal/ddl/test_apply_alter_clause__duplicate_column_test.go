//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterClause_DuplicateColumn 테스트
package ddl

import "testing"

func TestApplyAlterClause_DuplicateColumn(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "name", Raw: "name TEXT"}}}
	applyAlterClause(tbl, "ADD COLUMN name TEXT")
	if len(tbl.Columns) != 1 {
		t.Fatalf("expected 1 column (no duplicate), got %d", len(tbl.Columns))
	}
}
