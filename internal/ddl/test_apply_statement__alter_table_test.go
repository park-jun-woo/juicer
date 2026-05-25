//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyStatement_AlterTable 테스트
package ddl

import "testing"

func TestApplyStatement_AlterTable(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyStatement(tables, "ALTER TABLE users ADD COLUMN email TEXT")
	if len(tables["users"].Columns) != 2 {
		t.Fatal("expected 2 columns")
	}
}
