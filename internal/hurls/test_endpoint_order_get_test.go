//ff:func feature=hurl type=parse control=sequence
//ff:what TestEndpointOrder_GET 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestEndpointOrder_GET(t *testing.T) {
	ep := scanner.Endpoint{Method: "GET", Path: "/api/users"}
	if got := endpointOrder(ep); got != 1 {
		t.Fatalf("expected 1, got %d", got)
	}
}
