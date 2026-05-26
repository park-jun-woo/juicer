//ff:func feature=sql type=test control=sequence
//ff:what TestSqlcHint_InsertCov 테스트
package sqls

import "testing"

func TestSqlcHint_InsertCov(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "INSERT", Returns: []string{"error"}}
	got := sqlcHint(sk)
	if got == "" {
		t.Fatal("expected non-empty")
	}
}
