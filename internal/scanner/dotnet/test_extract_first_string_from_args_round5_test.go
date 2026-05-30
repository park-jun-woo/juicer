//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractFirstStringFromArgs_Round5 테스트
package dotnet

import "testing"

func TestExtractFirstStringFromArgs_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { F("a", "b"); } }`)
	args := firstOfType(t, root, "argument_list")
	if got := extractFirstStringFromArgs(args, src); got != "a" {
		t.Fatalf("got %q", got)
	}
	root2, src2 := parseCS(t, `class C { void M() { F(1, 2); } }`)
	args2 := firstOfType(t, root2, "argument_list")
	if got := extractFirstStringFromArgs(args2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
