//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineCRUDFromAST_Nil 테스트
package sqls

import "testing"

func TestRefineCRUDFromAST_Nil(t *testing.T) {
	got := refineCRUDFromAST(nil)
	if got != "EXEC" {
		t.Fatalf("expected EXEC, got %s", got)
	}
}
