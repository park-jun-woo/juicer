//ff:func feature=scan type=test control=sequence
//ff:what TestScanBody_NilBody 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanBody_NilBody(t *testing.T) {
	ep := &scanner.Endpoint{}
	scanBody(ep, nil, "c", nil, nil, "handler")
	if ep.Request != nil {
		t.Fatal("nil body should be a no-op")
	}
}
