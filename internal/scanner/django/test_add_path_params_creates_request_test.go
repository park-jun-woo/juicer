//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddPathParams_CreatesRequest 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddPathParams_CreatesRequest(t *testing.T) {
	ep := &scanner.Endpoint{}
	addPathParams(ep, []urlParam{{name: "id", converter: "int"}})
	if ep.Request == nil {
		t.Fatal("expected Request created")
	}
	if len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("unexpected path params: %+v", ep.Request.PathParams)
	}
}
