//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_AlterTableBranch 테스트
package ddl

import "testing"

func TestApplyStatement_AlterTableBranch(t *testing.T) {
	tables := map[string]*Table{"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}}}
	applyStatement(tables, "ALTER TABLE users ADD COLUMN email TEXT")
	if len(tables["users"].Columns) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(tables["users"].Columns))
	}
}
