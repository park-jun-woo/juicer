//ff:func feature=sql type=test control=sequence
//ff:what TestSqlcHint_Select 테스트
package sqls

import "testing"

func TestSqlcHint_Select(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "SELECT", Returns: []string{"*User"}}
	got := sqlcHint(sk)
	if got != ":one" {
		t.Fatalf("expected :one, got %s", got)
	}
}
