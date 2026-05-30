//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestAssignReturnTypeInfo 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAssignReturnTypeInfo(t *testing.T) {
	resp := &scanner.Response{}
	assignReturnTypeInfo(endpointInfo{returnType: "UserDto", returnIsArray: true}, resp)
	if resp.TypeName != "UserDto" || resp.Body != "array" {
		t.Fatalf("got %+v", resp)
	}
}
