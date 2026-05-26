//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUDIfNeeded_Select 테스트
package sqls

import "testing"

func TestRefineCRUDIfNeeded_Select(t *testing.T) {
	got := refineCRUDIfNeeded("SELECT", []string{"SELECT * FROM users"}, nil)
	if got != "SELECT" {
		t.Fatalf("expected SELECT, got %s", got)
	}
}

