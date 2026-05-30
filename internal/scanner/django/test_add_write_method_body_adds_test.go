//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddWriteMethodBody_Adds 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddWriteMethodBody_Adds(t *testing.T) {
	ep := &scanner.Endpoint{}
	serializers := map[string]serializerInfo{
		"S": {name: "User", fields: []scanner.Field{{Name: "id"}}},
	}
	addWriteMethodBody(ep, "POST", "S", serializers)
	if ep.Request == nil || ep.Request.Body == nil || ep.Request.Body.TypeName != "User" {
		t.Fatalf("expected body added, got %+v", ep.Request)
	}
}
