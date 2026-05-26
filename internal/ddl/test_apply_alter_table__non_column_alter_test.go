//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_NonColumnAlter ALTER COLUMN이 Raw만 수정하고 컬럼 수 유지 테스트
package ddl

import "testing"

func TestApplyAlterTable_NonColumnAlter(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyAlterTable(tables, "users", "ALTER COLUMN id SET NOT NULL")
	if len(tables["users"].Columns) != 1 {
		t.Fatal("columns count should be unchanged")
	}
	if tables["users"].Columns[0].Raw != "id INT NOT NULL" {
		t.Fatalf("expected Raw modified, got %q", tables["users"].Columns[0].Raw)
	}
}
