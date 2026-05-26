//ff:func feature=sql type=test control=sequence
//ff:what TestRefineExecCRUD_Insert 테스트
package sqls

import "testing"

func TestRefineExecCRUD_Insert(t *testing.T) {
	got := refineExecCRUD([]string{"INSERT INTO users"}, nil)
	if got != "INSERT" {
		t.Fatalf("expected INSERT, got %s", got)
	}
}
