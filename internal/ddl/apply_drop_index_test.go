//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyDropIndex_Existing 테스트
package ddl

import "testing"

func TestApplyDropIndex_Existing(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Indexes: []string{"CREATE INDEX idx_name ON users (name)"}},
	}
	applyDropIndex(tables, "idx_name")
	if len(tables["users"].Indexes) != 0 {
		t.Fatalf("expected 0 indexes, got %d", len(tables["users"].Indexes))
	}
}
