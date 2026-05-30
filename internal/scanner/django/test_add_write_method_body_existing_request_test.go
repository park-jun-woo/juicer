//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddWriteMethodBody_ExistingRequest 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddWriteMethodBody_ExistingRequest(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{}}
	addWriteMethodBody(ep, "PATCH", "S", map[string]serializerInfo{"S": {name: "T"}})
	if ep.Request.Body == nil || ep.Request.Body.TypeName != "T" {
		t.Fatalf("expected body set, got %+v", ep.Request.Body)
	}
}
