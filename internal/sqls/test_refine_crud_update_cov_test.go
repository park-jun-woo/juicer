//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUD_UpdateCov 테스트
package sqls

import "testing"

func TestRefineCRUD_UpdateCov(t *testing.T) {
	if refineCRUD([]string{"UPDATE users SET name = $1"}) != "UPDATE" {
		t.Fatal("expected UPDATE")
	}
}
