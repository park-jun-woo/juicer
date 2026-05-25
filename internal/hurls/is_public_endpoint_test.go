package hurls

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestIsPublicEndpoint_Health(t *testing.T) {
	ep := scanner.Endpoint{Method: "GET", Path: "/api/health"}
	if !isPublicEndpoint(ep) {
		t.Fatal("expected true")
	}
}

func TestIsPublicEndpoint_Login(t *testing.T) {
	ep := scanner.Endpoint{Method: "POST", Path: "/api/auth/login"}
	if !isPublicEndpoint(ep) {
		t.Fatal("expected true")
	}
}

func TestIsPublicEndpoint_SMS(t *testing.T) {
	ep := scanner.Endpoint{Method: "POST", Path: "/sms"}
	if !isPublicEndpoint(ep) {
		t.Fatal("expected true")
	}
}

func TestIsPublicEndpoint_Regular(t *testing.T) {
	ep := scanner.Endpoint{Method: "GET", Path: "/api/users"}
	if isPublicEndpoint(ep) {
		t.Fatal("expected false")
	}
}
