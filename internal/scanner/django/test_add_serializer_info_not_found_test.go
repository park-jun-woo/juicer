//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddSerializerInfo_NotFound 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddSerializerInfo_NotFound(t *testing.T) {
	ep := &scanner.Endpoint{}
	am := actionMethod{method: "POST"}
	addSerializerInfo(ep, am, "Missing", map[string]serializerInfo{})
	if ep.Request != nil {
		t.Fatal("expected no body when serializer not in map")
	}
}
