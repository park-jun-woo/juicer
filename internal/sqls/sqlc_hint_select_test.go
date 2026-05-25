package sqls

import "testing"

func TestSqlcHintSelect_One(t *testing.T) {
	sk := &MethodSkeleton{Returns: []string{"*User"}}
	got := sqlcHintSelect(sk)
	if got != ":one" {
		t.Fatalf("expected :one, got %s", got)
	}
}

func TestSqlcHintSelect_Many(t *testing.T) {
	sk := &MethodSkeleton{Returns: []string{"[]User"}}
	got := sqlcHintSelect(sk)
	if got != ":many" {
		t.Fatalf("expected :many, got %s", got)
	}
}
