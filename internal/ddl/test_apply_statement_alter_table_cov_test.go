//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_AlterTableCov 테스트
package ddl

import "testing"

func TestApplyStatement_AlterTableCov(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyStatement(tables, "ALTER TABLE users ADD COLUMN email TEXT")
	if len(tables["users"].Columns) != 2 {
		t.Fatal("expected 2 columns")
	}
}
