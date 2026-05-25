package sqls

import "testing"

func TestSqlcHint_Select(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "SELECT", Returns: []string{"*User"}}
	got := sqlcHint(sk)
	if got != ":one" {
		t.Fatalf("expected :one, got %s", got)
	}
}

func TestSqlcHint_SelectMany(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "SELECT", Returns: []string{"[]User"}}
	got := sqlcHint(sk)
	if got != ":many" {
		t.Fatalf("expected :many, got %s", got)
	}
}

func TestSqlcHint_Insert(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "INSERT"}
	got := sqlcHint(sk)
	if got != ":exec" {
		t.Fatalf("expected :exec, got %s", got)
	}
}

func TestSqlcHint_Delete(t *testing.T) {
	sk := &MethodSkeleton{CRUD: "DELETE"}
	got := sqlcHint(sk)
	if got != ":exec" {
		t.Fatalf("expected :exec, got %s", got)
	}
}
