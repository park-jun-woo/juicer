//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestAddSerializerInfo_EmptyClass 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAddSerializerInfo_EmptyClass(t *testing.T) {
	ep := &scanner.Endpoint{}
	addSerializerInfo(ep, actionMethod{method: "POST"}, "", nil)
	if ep.Request != nil {
		t.Fatal("expected no body for empty serializer class")
	}
}
