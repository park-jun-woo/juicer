//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineCRUDIfNeeded_Exec 테스트
package sqls

import "testing"

func TestRefineCRUDIfNeeded_Exec(t *testing.T) {
	got := refineCRUDIfNeeded("EXEC", []string{"INSERT INTO users"}, nil)
	if got != "INSERT" {
		t.Fatalf("expected INSERT, got %s", got)
	}
}
