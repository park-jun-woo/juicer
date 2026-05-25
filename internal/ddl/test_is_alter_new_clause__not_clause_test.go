//ff:func feature=ddl type=parse control=sequence
//ff:what TestIsAlterNewClause_NotClause 테스트
package ddl

import "testing"

func TestIsAlterNewClause_NotClause(t *testing.T) {
	if isAlterNewClause("name TEXT NOT NULL") {
		t.Fatal("expected false")
	}
}
