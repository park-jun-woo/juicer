//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_CreateIndexCov 테스트
package ddl

import "testing"

func TestApplyStatement_CreateIndexCov(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyStatement(tables, "CREATE INDEX idx_name ON users (name)")
	if len(tables["users"].Indexes) != 1 {
		t.Fatal("expected 1 index")
	}
}
