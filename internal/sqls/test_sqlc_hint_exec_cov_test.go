//ff:func feature=sql type=test control=sequence
//ff:what TestSqlcHint_ExecCov 테스트
package sqls

import "testing"

func TestSqlcHint_ExecCov(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "UPDATE"}
	got := sqlcHint(sk)
	if got != ":exec" {
		t.Fatalf("expected :exec, got %s", got)
	}
}
