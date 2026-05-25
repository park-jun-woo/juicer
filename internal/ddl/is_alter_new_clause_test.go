//ff:func feature=ddl type=parse control=sequence
//ff:what TestIsAlterNewClause_AddColumn 테스트
package ddl

import "testing"

func TestIsAlterNewClause_AddColumn(t *testing.T) {
	if !isAlterNewClause("ADD COLUMN name TEXT") {
		t.Fatal("expected true")
	}
}
