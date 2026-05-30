//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestEnsureRequest_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestEnsureRequest_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("request should be created")
	}
	prev := ep.Request
	ensureRequest(ep)
	if ep.Request != prev {
		t.Fatal("ensureRequest should not replace existing request")
	}
}
