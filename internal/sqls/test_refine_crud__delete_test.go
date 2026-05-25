//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineCRUD_Delete 테스트
package sqls

import "testing"

func TestRefineCRUD_Delete(t *testing.T) {
	got := refineCRUD([]string{"DELETE FROM users"})
	if got != "DELETE" {
		t.Fatalf("expected DELETE, got %s", got)
	}
}
