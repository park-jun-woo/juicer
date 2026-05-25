//ff:func feature=hurl type=parse control=sequence
//ff:what TestEndpointOrder_POST 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestEndpointOrder_POST(t *testing.T) {
	ep := scanner.Endpoint{Method: "POST", Path: "/api/users"}
	if got := endpointOrder(ep); got != 2 {
		t.Fatalf("expected 2, got %d", got)
	}
}
