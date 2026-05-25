package hurls

import "testing"

func TestMatchesEndpoint_Host(t *testing.T) {
	content := "GET {{host}}/api/health\nHTTP 200\n"
	if !matchesEndpoint(content, "GET", "/api/health") {
		t.Fatal("expected true")
	}
}

func TestMatchesEndpoint_HTTPS(t *testing.T) {
	content := "GET https://{{host}}/api/health\nHTTP 200\n"
	if !matchesEndpoint(content, "GET", "/api/health") {
		t.Fatal("expected true")
	}
}

func TestMatchesEndpoint_HTTP(t *testing.T) {
	content := "GET http://{{host}}/api/health\nHTTP 200\n"
	if !matchesEndpoint(content, "GET", "/api/health") {
		t.Fatal("expected true")
	}
}

func TestMatchesEndpoint_NoMatch(t *testing.T) {
	content := "GET {{host}}/api/users\nHTTP 200\n"
	if matchesEndpoint(content, "GET", "/api/health") {
		t.Fatal("expected false")
	}
}
