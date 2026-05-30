//ff:func feature=scan type=test control=sequence
//ff:what TestScanBody_NonCtxCall 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanBody_NonCtxCall(t *testing.T) {

	src := `package m
func h() {
	doWork()
	other.Method()
}
`
	body := handlerBody(t, src)
	ep := &scanner.Endpoint{}

	scanBody(ep, body, "c", nil, &funcIndex{}, "handler")
	if ep.Request != nil || len(ep.Responses) != 0 {
		t.Fatalf("non-ctx calls should not populate ep: %+v", ep)
	}
}
