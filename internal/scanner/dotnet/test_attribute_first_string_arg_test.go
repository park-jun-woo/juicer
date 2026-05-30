//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestAttributeFirstStringArg 테스트
package dotnet

import "testing"

func TestAttributeFirstStringArg(t *testing.T) {
	root, src := parseCS(t, `class C { [Route("api/users")] void m() {} }`)
	attrs := findAllByType(root, "attribute")
	if len(attrs) == 0 {
		t.Skip("no attribute")
	}
	if got := attributeFirstStringArg(attrs[0], src); got != "api/users" {
		t.Fatalf("got %q", got)
	}
}
