//ff:func feature=sql type=test control=sequence
//ff:what TestRefineExecCRUD_FromASTCov 테스트
package sqls

import "testing"

func TestRefineExecCRUD_FromASTCov(t *testing.T) {
	got := refineExecCRUD(nil, nil)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
