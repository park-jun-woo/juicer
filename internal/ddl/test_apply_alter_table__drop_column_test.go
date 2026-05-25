//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterTable_DropColumn 테스트
package ddl

import "testing"

func TestApplyAlterTable_DropColumn(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}, {Name: "email", Raw: "email TEXT"}}},
	}
	applyAlterTable(tables, "users", "DROP COLUMN email")
	if len(tables["users"].Columns) != 1 {
		t.Fatalf("expected 1 column, got %d", len(tables["users"].Columns))
	}
}
