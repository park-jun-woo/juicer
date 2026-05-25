//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyStatement_DropIndex 테스트
package ddl

import "testing"

func TestApplyStatement_DropIndex(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Indexes: []string{"CREATE INDEX idx_name ON users (name)"}},
	}
	applyStatement(tables, "DROP INDEX idx_name")
	if len(tables["users"].Indexes) != 0 {
		t.Fatal("expected 0 indexes")
	}
}
