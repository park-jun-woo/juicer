//ff:func feature=hurl type=parse control=sequence
//ff:what TestIsPublicEndpoint_Health 테스트
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
