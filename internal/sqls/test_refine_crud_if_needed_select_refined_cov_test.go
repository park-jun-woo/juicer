//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUDIfNeeded_SelectRefinedCov 테스트
package sqls

import "testing"

func TestRefineCRUDIfNeeded_SelectRefinedCov(t *testing.T) {
	got := refineCRUDIfNeeded("SELECT", []string{"DELETE FROM users"}, nil)
	if got != "DELETE" {
		t.Fatalf("expected DELETE, got %s", got)
	}
}
