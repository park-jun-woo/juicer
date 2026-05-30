//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestFindIntLiteralInArg_Round5 테스트
package dotnet

import "testing"

func TestFindIntLiteralInArg_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { F(42); } }`)
	arg := firstOfType(t, root, "argument")
	v, ok := findIntLiteralInArg(arg, src)
	if !ok || v != 42 {
		t.Fatalf("got %d %v", v, ok)
	}
	root2, src2 := parseCS(t, `class C { void M() { F("x"); } }`)
	arg2 := firstOfType(t, root2, "argument")
	if _, ok := findIntLiteralInArg(arg2, src2); ok {
		t.Fatal("expected no int literal")
	}
}
