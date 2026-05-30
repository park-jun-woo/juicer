//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestBuildBuilderEndpoint_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildBuilderEndpoint_Round5(t *testing.T) {
	br := builderRoute{method: "GET", path: "/users/{id}", handler: "get_user"}
	ep := buildBuilderEndpoint(br, structIndex{}, map[string][]scanner.Field{}, map[string]*handlerInfo{})
	if ep.Method != "GET" || ep.Path != "/users/{id}" {
		t.Fatalf("endpoint: %+v", ep)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("path params: %+v", ep.Request)
	}
}
