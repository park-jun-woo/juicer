//ff:func feature=scan type=test topic=flask control=sequence
//ff:what applyJSONBody hasJSONBody일 때 inline json body 설정 테스트
package flask

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyJSONBody(t *testing.T) {
	// no json body -> nothing
	ep := &scanner.Endpoint{}
	applyJSONBody(ep, routeInfo{})
	if ep.Request != nil {
		t.Error("no json -> no request")
	}
	// json body with fields
	ep2 := &scanner.Endpoint{}
	applyJSONBody(ep2, routeInfo{hasJSONBody: true, jsonFields: []string{"name", "email"}})
	if ep2.Request == nil || ep2.Request.Body == nil {
		t.Fatalf("body not set: %+v", ep2.Request)
	}
	if ep2.Request.Body.Method != "get_json" || len(ep2.Request.Body.Fields) != 2 {
		t.Errorf("body: %+v", ep2.Request.Body)
	}
}
