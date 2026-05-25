package hurls

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestBuildEndpointStatuses(t *testing.T) {
	eps := []scanner.Endpoint{
		{Method: "GET", Path: "/api/health"},
		{Method: "POST", Path: "/api/users"},
	}
	statuses := buildEndpointStatuses(eps)
	if len(statuses) != 2 {
		t.Fatalf("expected 2, got %d", len(statuses))
	}
	if statuses[0].ID != "GET /api/health" {
		t.Fatalf("expected GET /api/health, got %s", statuses[0].ID)
	}
	if statuses[0].Status != "TODO" {
		t.Fatalf("expected TODO, got %s", statuses[0].Status)
	}
}

func TestBuildEndpointStatuses_Empty(t *testing.T) {
	statuses := buildEndpointStatuses(nil)
	if len(statuses) != 0 {
		t.Fatalf("expected 0, got %d", len(statuses))
	}
}
