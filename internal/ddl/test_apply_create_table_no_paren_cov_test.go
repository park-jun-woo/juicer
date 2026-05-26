//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyCreateTable_NoParenCov 테스트
package ddl

import "testing"

func TestApplyCreateTable_NoParenCov(t *testing.T) {
	tables := make(map[string]*Table)
	applyCreateTable(tables, "empty", "CREATE TABLE empty")
	if tables["empty"] == nil {
		t.Fatal("expected table even without paren body")
	}
}
