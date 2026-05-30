//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestMethodLevelRoute_Round5 테스트
package dotnet

import "testing"

func TestMethodLevelRoute_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [Route("custom")] public void M() {} }`)
	m := firstOfType(t, root, "method_declaration")
	if got := methodLevelRoute(m, src); got != "custom" {
		t.Fatalf("got %q", got)
	}
	root2, src2 := parseCS(t, `class C { public void M() {} }`)
	m2 := firstOfType(t, root2, "method_declaration")
	if got := methodLevelRoute(m2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
