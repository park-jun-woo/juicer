//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyStatement_DropTable 테스트
package ddl

import "testing"

func TestApplyStatement_DropTable(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyStatement(tables, "DROP TABLE users")
	if _, ok := tables["users"]; ok {
		t.Fatal("expected users to be deleted")
	}
}
