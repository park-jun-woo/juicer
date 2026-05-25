//ff:func feature=ddl type=parse control=sequence
//ff:what TestIsAlterNewClause_DropColumn 테스트
package ddl

import "testing"

func TestIsAlterNewClause_DropColumn(t *testing.T) {
	if !isAlterNewClause("DROP COLUMN name") {
		t.Fatal("expected true")
	}
}
