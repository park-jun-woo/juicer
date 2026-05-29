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

	// ADD COLUMN with comment-only definition -> empty column name, not added
	applyAlterClause(tbl, "ADD COLUMN -- just a comment")
	if len(tbl.Columns) != 1 {
		t.Fatalf("comment-only ADD COLUMN should not add a column, got %d", len(tbl.Columns))
	}

	// ALTER COLUMN -> delegates to applyAlterColumn (mutates column Raw)
	applyAlterClause(tbl, "ALTER COLUMN id SET NOT NULL")
	if len(tbl.Columns) != 1 {
		t.Fatalf("ALTER COLUMN should not change column count, got %d", len(tbl.Columns))
	}

	// ADD CONSTRAINT -> appends a constraint
	applyAlterClause(tbl, "ADD CONSTRAINT uq_users_id UNIQUE (id)")
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint after ADD CONSTRAINT, got %d", len(tbl.Constraints))
	}

	// DROP CONSTRAINT -> removes the named constraint
	applyAlterClause(tbl, "DROP CONSTRAINT uq_users_id")
	if len(tbl.Constraints) != 0 {
		t.Fatalf("expected 0 constraints after DROP CONSTRAINT, got %d: %v", len(tbl.Constraints), tbl.Constraints)
	}

	// unrecognized clause (no-op)
	applyAlterClause(tbl, "RENAME TO new_users")
	if len(tbl.Columns) != 1 {
		t.Fatal("unrecognized clause should not change columns")
	}
}
