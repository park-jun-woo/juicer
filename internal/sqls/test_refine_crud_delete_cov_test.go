//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUD_DeleteCov 테스트
package sqls

import "testing"

func TestRefineCRUD_DeleteCov(t *testing.T) {
	if refineCRUD([]string{"DELETE FROM users"}) != "DELETE" {
		t.Fatal("expected DELETE")
	}
}
