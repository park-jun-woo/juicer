//ff:func feature=ddl type=test control=sequence
//ff:what applyAlterColumn의 컬럼 Raw 수정 및 미존재 컬럼 무시 테스트
package ddl

import "testing"

func TestApplyAlterColumn(t *testing.T) {
	tbl := &Table{
		Name: "users",
		Columns: []Column{
			{Name: "email", Raw: "email TEXT NOT NULL DEFAULT ''"},
			{Name: "role", Raw: "role TEXT NOT NULL"},
		},
	}

	// SET NOT NULL on role (already has it)
	applyAlterColumn(tbl, "role", "SET NOT NULL")
	if tbl.Columns[1].Raw != "role TEXT NOT NULL" {
		t.Fatalf("expected no change, got %q", tbl.Columns[1].Raw)
	}

	// DROP NOT NULL on email
	applyAlterColumn(tbl, "email", "DROP NOT NULL")
	if tbl.Columns[0].Raw != "email TEXT DEFAULT ''" {
		t.Fatalf("expected NOT NULL removed, got %q", tbl.Columns[0].Raw)
	}

	// SET DEFAULT on role
	applyAlterColumn(tbl, "role", "SET DEFAULT 'admin'")
	if tbl.Columns[1].Raw != "role TEXT NOT NULL DEFAULT 'admin'" {
		t.Fatalf("expected DEFAULT added, got %q", tbl.Columns[1].Raw)
	}

	// Non-existent column (should be no-op)
	applyAlterColumn(tbl, "nonexistent", "SET NOT NULL")
	if len(tbl.Columns) != 2 {
		t.Fatal("columns count should be unchanged")
	}
}
