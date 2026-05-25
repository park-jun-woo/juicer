//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyCreateTable_NoParen 테스트
package ddl

import "testing"

func TestApplyCreateTable_NoParen(t *testing.T) {
	tables := make(map[string]*Table)
	applyCreateTable(tables, "empty", "CREATE TABLE empty")
	if tables["empty"] == nil {
		t.Fatal("expected table even without parens")
	}
}
