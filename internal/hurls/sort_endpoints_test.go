package hurls

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestSortEndpoints(t *testing.T) {
	eps := []scanner.Endpoint{
		{Method: "PUT", Path: "/api/users/:id"},
		{Method: "GET", Path: "/api/health"},
		{Method: "POST", Path: "/api/users"},
		{Method: "GET", Path: "/api/users"},
	}
	sorted := sortEndpoints(eps)
	if sorted[0].Path != "/api/health" {
		t.Fatalf("expected health first, got %s", sorted[0].Path)
	}
}

func TestSortEndpoints_Empty(t *testing.T) {
	sorted := sortEndpoints(nil)
	if len(sorted) != 0 {
		t.Fatalf("expected 0, got %d", len(sorted))
	}
}
