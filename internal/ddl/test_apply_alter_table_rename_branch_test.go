//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_RenameBranch 테스트
package ddl

import "testing"

func TestApplyAlterTable_RenameBranch(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyAlterTable(tables, "users", "RENAME TO people")
	if _, ok := tables["people"]; !ok {
		t.Fatal("expected renamed table 'people'")
	}
}
