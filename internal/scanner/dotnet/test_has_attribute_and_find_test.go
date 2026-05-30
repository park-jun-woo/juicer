//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestHasAttributeAndFind 테스트
package dotnet

import "testing"

func TestHasAttributeAndFind(t *testing.T) {
	root, src := parseCS(t, `[ApiController] class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if !hasAttribute(cls, src, "ApiController") {
		t.Fatal("ApiController")
	}
	if hasAttribute(cls, src, "Missing") {
		t.Fatal("missing")
	}
	if findAttribute(cls, src, "ApiController") == nil {
		t.Fatal("find")
	}
}
