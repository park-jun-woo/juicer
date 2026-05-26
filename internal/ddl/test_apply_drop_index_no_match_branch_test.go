//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyDropIndex_NoMatchBranch 테스트
package ddl

import "testing"

func TestApplyDropIndex_NoMatchBranch(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Indexes: []string{"CREATE INDEX idx_name ON users (name)"}},
	}
	applyDropIndex(tables, "idx_other")
	if len(tables["users"].Indexes) != 1 {
		t.Fatalf("expected 1 index, got %d", len(tables["users"].Indexes))
	}
}
