//ff:func feature=sql type=parse control=sequence
//ff:what TestSqlcHint_Insert 테스트
package sqls

import "testing"

func TestSqlcHint_Insert(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "INSERT"}
	got := sqlcHint(sk)
	if got != ":exec" {
		t.Fatalf("expected :exec, got %s", got)
	}
}
