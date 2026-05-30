//ff:func feature=scan type=test control=sequence topic=actix
//ff:what prefixEndpointsForHandler — 핸들러 일치 + 미prefix 경로에만 prefix 부여를 검증
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestPrefixEndpointsForHandler(t *testing.T) {
	eps := []scanner.Endpoint{
		{Handler: "list_users", Path: "/users"},       // matches, gets prefix
		{Handler: "other", Path: "/things"},           // handler mismatch, unchanged
		{Handler: "list_users", Path: "/api/already"}, // already prefixed, unchanged
	}
	prefixEndpointsForHandler(eps, "list_users", "/api")

	if eps[0].Path != "/api/users" {
		t.Errorf("eps[0].Path = %q, want /api/users", eps[0].Path)
	}
	if eps[1].Path != "/things" {
		t.Errorf("eps[1].Path = %q, want /things (unchanged)", eps[1].Path)
	}
	if eps[2].Path != "/api/already" {
		t.Errorf("eps[2].Path = %q, want /api/already (unchanged)", eps[2].Path)
	}
}
