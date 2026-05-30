//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestLastIdentifier 테스트
package dotnet

import "testing"

func TestLastIdentifier(t *testing.T) {
	root, src := parseCS(t, `class C { [System.Foo] void m() {} }`)
	qn := findAllByType(root, "qualified_name")
	if len(qn) == 0 {
		t.Skip("no qualified name")
	}
	if got := lastIdentifier(qn[0], src); got != "Foo" {
		t.Fatalf("got %q", got)
	}
}
