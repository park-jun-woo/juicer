//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddSerializerInfo_Adds 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddSerializerInfo_Adds(t *testing.T) {
	ep := &scanner.Endpoint{}
	am := actionMethod{action: "create", method: "POST"}
	serializers := map[string]serializerInfo{
		"UserSerializer": {name: "User", fields: []scanner.Field{{Name: "id"}}},
	}
	addSerializerInfo(ep, am, "UserSerializer", serializers)
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected body added")
	}
	if ep.Request.Body.TypeName != "User" {
		t.Errorf("TypeName = %q, want User", ep.Request.Body.TypeName)
	}
}
