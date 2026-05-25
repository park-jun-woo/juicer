//ff:func feature=hurl type=parse control=sequence
//ff:what TestEndpointOrder_PUT 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestEndpointOrder_PUT(t *testing.T) {
	ep := scanner.Endpoint{Method: "PUT", Path: "/api/users"}
	if got := endpointOrder(ep); got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}
