//ff:func feature=sql type=parse control=sequence
//ff:what TestSqlcHint_Delete 테스트
package sqls

import "testing"

func TestSqlcHint_Delete(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "DELETE"}
	got := sqlcHint(sk)
	if got != ":exec" {
		t.Fatalf("expected :exec, got %s", got)
	}
}
