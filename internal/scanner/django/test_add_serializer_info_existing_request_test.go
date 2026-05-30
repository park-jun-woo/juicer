//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddSerializerInfo_ExistingRequest 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddSerializerInfo_ExistingRequest(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{}}
	am := actionMethod{method: "PUT"}
	addSerializerInfo(ep, am, "S", map[string]serializerInfo{"S": {name: "T"}})
	if ep.Request.Body == nil || ep.Request.Body.TypeName != "T" {
		t.Fatalf("expected body set on existing request, got %+v", ep.Request.Body)
	}
}
