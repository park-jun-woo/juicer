package scanner

import "testing"

func TestJoinPath_Both(t *testing.T) {
	got := joinPath("/api", "/v1")
	if got != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s", got)
	}
}

func TestJoinPath_EmptyA(t *testing.T) {
	got := joinPath("", "/v1")
	if got != "/v1" {
		t.Fatalf("expected /v1, got %s", got)
	}
}
