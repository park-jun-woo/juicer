//ff:func feature=hurl type=parse control=sequence
//ff:what TestFindEndpoint_NotFound 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestFindEndpoint_NotFound(t *testing.T) {
	eps := []scanner.Endpoint{{Method: "GET", Path: "/health"}}
	ep := findEndpoint(eps, "DELETE /users")
	if ep != nil {
		t.Fatal("expected nil")
	}
}
