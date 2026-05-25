//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterClause_Empty 테스트
package ddl

import "testing"

func TestApplyAlterClause_Empty(t *testing.T) {
	tbl := &Table{Name: "users"}
	applyAlterClause(tbl, "")
}
