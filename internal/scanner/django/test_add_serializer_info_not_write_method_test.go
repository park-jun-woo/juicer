//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddSerializerInfo_NotWriteMethod 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddSerializerInfo_NotWriteMethod(t *testing.T) {
	ep := &scanner.Endpoint{}
	am := actionMethod{action: "list", method: "GET"}
	addSerializerInfo(ep, am, "UserSerializer", map[string]serializerInfo{
		"UserSerializer": {name: "User"},
	})
	if ep.Request != nil {
		t.Fatal("expected no body for read method")
	}
}
