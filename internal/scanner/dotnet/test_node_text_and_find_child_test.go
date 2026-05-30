//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestNodeTextAndFindChild 테스트
package dotnet

import "testing"

func TestNodeTextAndFindChild(t *testing.T) {
	root, src := parseCS(t, `class C {}`)
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no class")
	}
	id := findChildByType(classes[0], "identifier")
	if id != nil && nodeText(id, src) != "C" {
		t.Fatalf("name %q", nodeText(id, src))
	}
	if findChildByType(classes[0], "nonexistent") != nil {
		t.Fatal("nil expected")
	}
}
