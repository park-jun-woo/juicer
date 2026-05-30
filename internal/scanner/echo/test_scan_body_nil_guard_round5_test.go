//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestScanBody_NilGuard_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanBody_NilGuard_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}

	scanBody(ep, nil, "c", nil, buildFuncIndex(nil), "handler")
}
