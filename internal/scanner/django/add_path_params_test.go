//ff:func feature=scan type=test control=sequence topic=django
//ff:what addPathParams — path 파라미터 추가 분기를 검증
package django

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAddPathParams_Empty(t *testing.T) {
	ep := &scanner.Endpoint{}
	addPathParams(ep, nil)
	if ep.Request != nil {
		t.Fatal("expected Request to stay nil for empty params")
	}
}

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

func TestAddPathParams_ExistingRequest(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{
		PathParams: []scanner.Param{{Name: "a"}},
	}}
	addPathParams(ep, []urlParam{{name: "b", converter: "str"}})
	if len(ep.Request.PathParams) != 2 {
		t.Fatalf("expected 2 path params, got %d", len(ep.Request.PathParams))
	}
}
