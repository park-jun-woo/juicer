//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_DropTableBranch 테스트
package ddl

import "testing"

func TestApplyStatement_DropTableBranch(t *testing.T) {
	tables := map[string]*Table{"users": {Name: "users"}}
	applyStatement(tables, "DROP TABLE users")
	if tables["users"] != nil {
		t.Fatal("expected users table to be dropped")
	}
}
