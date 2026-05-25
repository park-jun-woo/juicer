//ff:func feature=hurl type=parse control=sequence
//ff:what TestEndpointOrder_Public 테스트
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
