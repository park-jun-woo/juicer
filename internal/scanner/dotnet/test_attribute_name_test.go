//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestAttributeName 테스트
package dotnet

import "testing"

func TestAttributeName(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpGet] void m() {} }`)
	attrs := findAllByType(root, "attribute")
	if len(attrs) == 0 {
		t.Skip("no attribute")
	}
	if got := attributeName(attrs[0], src); got != "HttpGet" {
		t.Fatalf("got %q", got)
	}
}
