//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineExecCRUD_NoSQL 테스트
package sqls

import "testing"

func TestRefineExecCRUD_NoSQL(t *testing.T) {
	got := refineExecCRUD(nil, nil)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
