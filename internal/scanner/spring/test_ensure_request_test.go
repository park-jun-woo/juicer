//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestEnsureRequest 테스트
package spring

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
}
