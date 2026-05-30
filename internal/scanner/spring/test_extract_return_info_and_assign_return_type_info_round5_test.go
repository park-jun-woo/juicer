//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractReturnInfo_And_AssignReturnTypeInfo_Round5 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestExtractReturnInfo_And_AssignReturnTypeInfo_Round5(t *testing.T) {
	root, src := sParse(t, `class C { public java.util.List<UserDto> all() { return null; } }`)
	m := sFirst(t, root, "method_declaration")
	var ep endpointInfo
	extractReturnInfo(m, src, &ep)
	if ep.returnType == "" {
		t.Fatalf("returnType empty: %+v", ep)
	}
	var resp scanner.Response
	assignReturnTypeInfo(ep, &resp)
	if resp.TypeName == "" {
		t.Fatalf("resp typename empty: %+v", resp)
	}
}
