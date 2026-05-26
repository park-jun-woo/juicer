//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUD_ExecCov 테스트
package sqls

import "testing"

func TestRefineCRUD_ExecCov(t *testing.T) {
	if refineCRUD([]string{"CALL some_proc()"}) != "EXEC" {
		t.Fatal("expected EXEC")
	}
}
