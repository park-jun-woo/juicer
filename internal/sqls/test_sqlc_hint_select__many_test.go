//ff:func feature=sql type=parse control=sequence
//ff:what TestSqlcHintSelect_Many 테스트
package sqls

import "testing"

func TestSqlcHintSelect_Many(t *testing.T) {
	sk := &MethodSkeleton{Returns: []string{"[]User"}}
	got := sqlcHintSelect(sk)
	if got != ":many" {
		t.Fatalf("expected :many, got %s", got)
	}
}
