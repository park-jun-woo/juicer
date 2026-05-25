package hurls

import "testing"

func TestParseEndpointID(t *testing.T) {
	method, path := parseEndpointID("GET /api/health")
	if method != "GET" || path != "/api/health" {
		t.Fatalf("got %q, %q", method, path)
	}
}

func TestParseEndpointID_NoSpace(t *testing.T) {
	method, path := parseEndpointID("GET")
	if method != "GET" || path != "" {
		t.Fatalf("got %q, %q", method, path)
	}
}
