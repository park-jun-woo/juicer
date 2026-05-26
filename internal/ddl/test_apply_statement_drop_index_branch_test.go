//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_DropIndexBranch 테스트
package ddl

import "testing"

func TestApplyStatement_DropIndexBranch(t *testing.T) {
	tables := map[string]*Table{"users": {Name: "users", Indexes: []string{"CREATE INDEX idx_name ON users (name)"}}}
	applyStatement(tables, "DROP INDEX idx_name")
}
