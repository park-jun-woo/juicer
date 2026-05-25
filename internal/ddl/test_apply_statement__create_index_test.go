//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyStatement_CreateIndex 테스트
package ddl

import "testing"

func TestApplyStatement_CreateIndex(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyStatement(tables, "CREATE INDEX idx_name ON users (name)")
	if len(tables["users"].Indexes) != 1 {
		t.Fatal("expected 1 index")
	}
}
