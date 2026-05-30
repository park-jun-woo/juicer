//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestNodeTextFindChild 테스트
package spring

import "testing"

func TestNodeTextFindChild(t *testing.T) {
	root, src := parseS(t, `class C {}`)
	classes := findAllByType(root, "class_declaration")
	if findChildByType(classes[0], "class_body") == nil {
		t.Fatal("class_body")
	}
	if findChildByType(classes[0], "nope") != nil {
		t.Fatal("nil expected")
	}
	id := findChildByType(classes[0], "identifier")
	if id != nil && nodeText(id, src) != "C" {
		t.Fatalf("name %q", nodeText(id, src))
	}
}
