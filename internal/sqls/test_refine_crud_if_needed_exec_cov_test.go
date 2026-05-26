//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUDIfNeeded_ExecCov 테스트
package sqls

import "testing"

func TestRefineCRUDIfNeeded_ExecCov(t *testing.T) {
	got := refineCRUDIfNeeded("EXEC", []string{"INSERT INTO users"}, nil)
	if got != "INSERT" {
		t.Fatalf("expected INSERT, got %s", got)
	}
}
