//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestEnsureRequest 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestEnsureRequest(t *testing.T) {
	ep := &scanner.Endpoint{}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	prev := ep.Request
	ensureRequest(ep)
	if ep.Request != prev {
		t.Fatal("should not replace existing request")
	}
}
