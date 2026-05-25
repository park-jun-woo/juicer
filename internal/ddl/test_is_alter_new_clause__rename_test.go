//ff:func feature=ddl type=parse control=sequence
//ff:what TestIsAlterNewClause_Rename 테스트
package ddl

import "testing"

func TestIsAlterNewClause_Rename(t *testing.T) {
	if !isAlterNewClause("RENAME TO new_name") {
		t.Fatal("expected true")
	}
}
