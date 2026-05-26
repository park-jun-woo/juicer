//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterClause_DropColumnCov 테스트
package ddl

import "testing"

func TestApplyAlterClause_DropColumnCov(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}, {Name: "name", Raw: "name TEXT"}}}
	applyAlterClause(tbl, "DROP COLUMN name")
	if len(tbl.Columns) != 1 {
		t.Fatalf("expected 1 column, got %d", len(tbl.Columns))
	}
}
