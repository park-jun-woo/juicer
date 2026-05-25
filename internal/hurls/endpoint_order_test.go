package hurls

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestEndpointOrder_Public(t *testing.T) {
	ep := scanner.Endpoint{Method: "GET", Path: "/api/health"}
	if got := endpointOrder(ep); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestEndpointOrder_GET(t *testing.T) {
	ep := scanner.Endpoint{Method: "GET", Path: "/api/users"}
	if got := endpointOrder(ep); got != 1 {
		t.Fatalf("expected 1, got %d", got)
	}
}

func TestEndpointOrder_POST(t *testing.T) {
	ep := scanner.Endpoint{Method: "POST", Path: "/api/users"}
	if got := endpointOrder(ep); got != 2 {
		t.Fatalf("expected 2, got %d", got)
	}
}

func TestEndpointOrder_PUT(t *testing.T) {
	ep := scanner.Endpoint{Method: "PUT", Path: "/api/users"}
	if got := endpointOrder(ep); got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}
