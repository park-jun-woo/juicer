package sqls

import "testing"

func TestSqlcHintInsert_NoReturning(t *testing.T) {
	sk := &MethodSkeleton{SQLFragments: []string{"INSERT INTO users"}}
	got := sqlcHintInsert(sk)
	if got != ":exec" {
		t.Fatalf("expected :exec, got %s", got)
	}
}

func TestSqlcHintInsert_WithReturning(t *testing.T) {
	sk := &MethodSkeleton{SQLFragments: []string{"INSERT INTO users RETURNING id"}}
	got := sqlcHintInsert(sk)
	if got != ":one" {
		t.Fatalf("expected :one, got %s", got)
	}
}
