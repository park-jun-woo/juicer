//ff:func feature=ddl type=parse control=sequence
//ff:what TestIsAlterNewClause_AlterColumn 테스트
package ddl

import "testing"

func TestIsAlterNewClause_AlterColumn(t *testing.T) {
	if !isAlterNewClause("ALTER COLUMN name SET NOT NULL") {
		t.Fatal("expected true")
	}
}
