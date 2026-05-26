//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUD_Insert 테스트
package sqls

import "testing"

func TestRefineCRUD_Insert(t *testing.T) {
	got := refineCRUD([]string{"INSERT INTO users"})
	if got != "INSERT" {
		t.Fatalf("expected INSERT, got %s", got)
	}
}
