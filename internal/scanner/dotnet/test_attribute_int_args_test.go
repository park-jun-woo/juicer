//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestAttributeIntArgs 테스트
package dotnet

import "testing"

func TestAttributeIntArgs(t *testing.T) {
	root, src := parseCS(t, `class C { [ProducesResponseType(201)] void m() {} }`)
	attrs := findAllByType(root, "attribute")
	if len(attrs) == 0 {
		t.Skip("no attribute")
	}
	ints := attributeIntArgs(attrs[0], src)
	if len(ints) != 1 || ints[0] != 201 {
		t.Fatalf("got %v", ints)
	}
}
