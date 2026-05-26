//ff:func feature=ddl type=test control=sequence
//ff:what applyAlterClause ADD/DROP/empty 분기 테스트
package ddl

import "testing"

func TestApplyAlterClause_Branches(t *testing.T) {
	tbl := &Table{Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}}

	// empty clause
	applyAlterClause(tbl, "")
	if len(tbl.Columns) != 1 {
		t.Fatal("empty clause should not change columns")
	}

	// ADD COLUMN
	applyAlterClause(tbl, "ADD COLUMN name TEXT")
	if len(tbl.Columns) != 2 {
		t.Fatal("expected 2 columns after ADD COLUMN")
	}

	// ADD COLUMN duplicate (should not add)
	applyAlterClause(tbl, "ADD COLUMN name TEXT")
	if len(tbl.Columns) != 2 {
		t.Fatal("duplicate ADD COLUMN should not add")
	}

	// DROP COLUMN
	applyAlterClause(tbl, "DROP COLUMN name")
	if len(tbl.Columns) != 1 {
		t.Fatal("expected 1 column after DROP COLUMN")
	}

	// unrecognized clause (no-op)
	applyAlterClause(tbl, "RENAME TO new_users")
	if len(tbl.Columns) != 1 {
		t.Fatal("unrecognized clause should not change columns")
	}
}
