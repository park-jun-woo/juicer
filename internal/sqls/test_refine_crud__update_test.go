//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineCRUD_Update 테스트
package sqls

import "testing"

func TestRefineCRUD_Update(t *testing.T) {
	got := refineCRUD([]string{"UPDATE users SET name = $1"})
	if got != "UPDATE" {
		t.Fatalf("expected UPDATE, got %s", got)
	}
}
