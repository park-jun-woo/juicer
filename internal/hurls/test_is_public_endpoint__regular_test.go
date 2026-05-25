//ff:func feature=hurl type=parse control=sequence
//ff:what TestIsPublicEndpoint_Regular 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestIsPublicEndpoint_Regular(t *testing.T) {
	ep := scanner.Endpoint{Method: "GET", Path: "/api/users"}
	if isPublicEndpoint(ep) {
		t.Fatal("expected false")
	}
}
