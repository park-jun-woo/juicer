//ff:func feature=sql type=parse control=sequence
//ff:what TestSqlcHint_SelectMany 테스트
package sqls

import "testing"

func TestSqlcHint_SelectMany(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "SELECT", Returns: []string{"[]User"}}
	got := sqlcHint(sk)
	if got != ":many" {
		t.Fatalf("expected :many, got %s", got)
	}
}
