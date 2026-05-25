//ff:func feature=sql type=parse control=sequence
//ff:what TestSqlcHintInsert_WithReturning 테스트
package sqls

import "testing"

func TestSqlcHintInsert_WithReturning(t *testing.T) {
	sk := &MethodSkeleton{SQLFragments: []string{"INSERT INTO users RETURNING id"}}
	got := sqlcHintInsert(sk)
	if got != ":one" {
		t.Fatalf("expected :one, got %s", got)
	}
}
