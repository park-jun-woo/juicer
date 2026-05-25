//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineCRUD_Default 테스트
package sqls

import "testing"

func TestRefineCRUD_Default(t *testing.T) {
	got := refineCRUD([]string{"CALL proc()"})
	if got != "EXEC" {
		t.Fatalf("expected EXEC, got %s", got)
	}
}
