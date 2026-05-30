//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddWriteMethodBody_NotFound 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddWriteMethodBody_NotFound(t *testing.T) {
	ep := &scanner.Endpoint{}
	addWriteMethodBody(ep, "PUT", "Missing", map[string]serializerInfo{})
	if ep.Request != nil {
		t.Fatal("expected no body when serializer missing")
	}
}
