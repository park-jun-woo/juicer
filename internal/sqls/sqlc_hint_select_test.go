//ff:func feature=sql type=test control=sequence
//ff:what TestSqlcHintSelect_One 테스트
package sqls

import "testing"

func TestSqlcHintSelect_One(t *testing.T) {
	sk := &MethodSkeleton{Returns: []string{"*User"}}
	got := sqlcHintSelect(sk)
	if got != ":one" {
		t.Fatalf("expected :one, got %s", got)
	}
}

