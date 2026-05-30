//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestFindStringLiteralInArg_Round5 테스트
package dotnet

import "testing"

func TestFindStringLiteralInArg_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { F("hello"); } }`)
	arg := firstOfType(t, root, "argument")
	if got := findStringLiteralInArg(arg, src); got != "hello" {
		t.Fatalf("got %q", got)
	}

	root2, src2 := parseCS(t, `class C { void M() { F(42); } }`)
	arg2 := firstOfType(t, root2, "argument")
	if got := findStringLiteralInArg(arg2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
