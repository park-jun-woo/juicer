//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestAttributeNamedArg_Round5 테스트
package dotnet

import "testing"

func TestAttributeNamedArg_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [Authorize(Roles = "Admin")] void M() {} }`)
	attr := firstOfType(t, root, "attribute")
	if got := attributeNamedArg(attr, src, "Roles"); got != "Admin" {
		t.Fatalf("got %q", got)
	}
	if got := attributeNamedArg(attr, src, "Policy"); got != "" {
		t.Fatalf("expected empty for missing named arg, got %q", got)
	}
}
