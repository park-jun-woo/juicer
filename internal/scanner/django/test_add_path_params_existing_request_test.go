//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddPathParams_ExistingRequest 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddPathParams_ExistingRequest(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{
		PathParams: []scanner.Param{{Name: "a"}},
	}}
	addPathParams(ep, []urlParam{{name: "b", converter: "str"}})
	if len(ep.Request.PathParams) != 2 {
		t.Fatalf("expected 2 path params, got %d", len(ep.Request.PathParams))
	}
}
