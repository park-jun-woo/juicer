//ff:func feature=scan type=test control=sequence
//ff:what TestScanBody_ChainedBind 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanBody_ChainedBind(t *testing.T) {
	src := `package m
func h() {
	c.Bind().Body(req)
}
`
	body := handlerBody(t, src)
	ep := &scanner.Endpoint{}
	scanBody(ep, body, "c", nil, &funcIndex{}, "handler")
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("chained Bind().Body should populate body: %+v", ep.Request)
	}
}
