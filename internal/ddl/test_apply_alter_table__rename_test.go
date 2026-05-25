//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterTable_Rename 테스트
package ddl

import "testing"

func TestApplyAlterTable_Rename(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyAlterTable(tables, "users", "RENAME TO accounts")
	if _, ok := tables["accounts"]; !ok {
		t.Fatal("expected table renamed to accounts")
	}
	if _, ok := tables["users"]; ok {
		t.Fatal("expected old name deleted")
	}
}
