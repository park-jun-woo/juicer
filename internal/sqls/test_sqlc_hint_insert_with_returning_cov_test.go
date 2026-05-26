//ff:func feature=sql type=test control=sequence
//ff:what TestSqlcHintInsert_WithReturningCov 테스트
package sqls

import "testing"

func TestSqlcHintInsert_WithReturningCov(t *testing.T) {
	sk := &MethodSkeleton{SQLFragments: []string{"INSERT INTO users RETURNING id"}}
	got := sqlcHintInsert(sk)
	if got != ":one" {
		t.Fatalf("expected :one, got %s", got)
	}
}
