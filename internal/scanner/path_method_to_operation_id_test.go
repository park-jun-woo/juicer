package scanner

import "testing"

func TestPathMethodToOperationID_Basic(t *testing.T) {
	got := pathMethodToOperationID("GET", "/api/v1/users")
	if got != "get_users" {
		t.Fatalf("expected get_users, got %s", got)
	}
}

func TestPathMethodToOperationID_WithParam(t *testing.T) {
	got := pathMethodToOperationID("GET", "/api/v1/users/:id")
	if got != "get_users_id" {
		t.Fatalf("expected get_users_id, got %s", got)
	}
}
