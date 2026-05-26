//ff:func feature=sql type=test control=sequence
//ff:what TestSqlcHintSelect_ManyCov 테스트
package sqls

import "testing"

func TestSqlcHintSelect_ManyCov(t *testing.T) {
	sk := &MethodSkeleton{Returns: []string{"[]User"}}
	got := sqlcHintSelect(sk)
	if got != ":many" {
		t.Fatalf("expected :many, got %s", got)
	}
}
