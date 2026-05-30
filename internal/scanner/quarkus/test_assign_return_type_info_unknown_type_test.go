//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAssignReturnTypeInfo_UnknownType 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAssignReturnTypeInfo_UnknownType(t *testing.T) {
	ep := endpointInfo{returnType: "SomeDto"}
	resp := &scanner.Response{}
	assignReturnTypeInfo(ep, resp)

	if resp.TypeName != "SomeDto" {
		t.Fatalf("got %+v", resp)
	}
}
