//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestEnsureRequest 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestEnsureRequest(t *testing.T) {
	ep := &scanner.Endpoint{Method: "GET"}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	existing := ep.Request
	ensureRequest(ep)
	if ep.Request != existing {
		t.Fatal("should not replace existing Request")
	}
}
