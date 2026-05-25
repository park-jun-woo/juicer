//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyDropIndex_NoMatch 테스트
package ddl

import "testing"

func TestApplyDropIndex_NoMatch(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Indexes: []string{"CREATE INDEX idx_name ON users (name)"}},
	}
	applyDropIndex(tables, "idx_other")
	if len(tables["users"].Indexes) != 1 {
		t.Fatalf("expected 1 index, got %d", len(tables["users"].Indexes))
	}
}
