//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestHandleQueryParam_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHandleQueryParam_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callWithStringArg(t, `"page"`)
	handleQueryParam(ep, call, "GET")
	if ep.Request == nil || len(ep.Request.Query) != 1 || ep.Request.Query[0].Name != "page" {
		t.Fatalf("query: %+v", ep.Request)
	}

	handleQueryParam(ep, call, "GET")
	if len(ep.Request.Query) != 1 {
		t.Fatalf("duplicate query: %+v", ep.Request.Query)
	}
}
