//ff:func feature=sql type=test control=sequence
//ff:what TestSqlcHintInsert_NoReturning 테스트
package sqls

import "testing"

func TestSqlcHintInsert_NoReturning(t *testing.T) {
	sk := &MethodSkeleton{SQLFragments: []string{"INSERT INTO users"}}
	got := sqlcHintInsert(sk)
	if got != ":exec" {
		t.Fatalf("expected :exec, got %s", got)
	}
}
