//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterTable_NonColumnAlter 테스트
package ddl

import "testing"

func TestApplyAlterTable_NonColumnAlter(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyAlterTable(tables, "users", "ALTER COLUMN id SET NOT NULL")
	if len(tables["users"].Columns) != 1 {
		t.Fatal("columns should be unchanged")
	}
}
