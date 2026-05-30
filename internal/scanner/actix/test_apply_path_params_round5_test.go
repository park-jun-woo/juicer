//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyPathParams_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyPathParams_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyPathParams(ep, "/users/{id}")
	if ep.Request == nil || len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("path params: %+v", ep.Request)
	}

	ep2 := &scanner.Endpoint{}
	applyPathParams(ep2, "/static")
	if ep2.Request != nil {
		t.Fatalf("expected nil request, got %+v", ep2.Request)
	}
}
