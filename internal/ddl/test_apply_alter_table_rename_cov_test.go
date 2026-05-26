//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_RenameCov 테스트
package ddl

import "testing"

func TestApplyAlterTable_RenameCov(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyAlterTable(tables, "users", "RENAME TO accounts")
	if tables["accounts"] == nil {
		t.Fatal("expected accounts table after rename")
	}
	if tables["users"] != nil {
		t.Fatal("expected users table to be removed")
	}
}
