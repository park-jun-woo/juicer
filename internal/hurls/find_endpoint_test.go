package hurls

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestFindEndpoint_Found(t *testing.T) {
	eps := []scanner.Endpoint{
		{Method: "GET", Path: "/health"},
		{Method: "POST", Path: "/users"},
	}
	ep := findEndpoint(eps, "POST /users")
	if ep == nil {
		t.Fatal("expected to find endpoint")
	}
}

func TestFindEndpoint_NotFound(t *testing.T) {
	eps := []scanner.Endpoint{{Method: "GET", Path: "/health"}}
	ep := findEndpoint(eps, "DELETE /users")
	if ep != nil {
		t.Fatal("expected nil")
	}
}
