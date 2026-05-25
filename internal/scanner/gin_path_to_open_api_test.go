package scanner

import "testing"

func TestGinPathToOpenAPI_Param(t *testing.T) {
	got := ginPathToOpenAPI("/api/users/:id")
	if got != "/api/users/{id}" {
		t.Fatalf("expected /api/users/{id}, got %s", got)
	}
}

func TestGinPathToOpenAPI_Wildcard(t *testing.T) {
	got := ginPathToOpenAPI("/files/*path")
	if got != "/files/{path}" {
		t.Fatalf("expected /files/{path}, got %s", got)
	}
}

func TestGinPathToOpenAPI_NoParam(t *testing.T) {
	got := ginPathToOpenAPI("/api/users")
	if got != "/api/users" {
		t.Fatalf("expected /api/users, got %s", got)
	}
}
