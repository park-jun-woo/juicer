//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddWriteMethodBody_NotWrite 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddWriteMethodBody_NotWrite(t *testing.T) {
	ep := &scanner.Endpoint{}
	addWriteMethodBody(ep, "GET", "S", map[string]serializerInfo{"S": {name: "User"}})
	if ep.Request != nil {
		t.Fatal("expected no body for GET")
	}
}
