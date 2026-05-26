//ff:func feature=sql type=test control=sequence
//ff:what TestRefineExecCRUD_FallbackExecCov 테스트
package sqls

import "testing"

func TestRefineExecCRUD_FallbackExecCov(t *testing.T) {
	got := refineExecCRUD([]string{"CALL proc()"}, nil)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
