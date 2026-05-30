//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestNodeTextAndFindChild 테스트
package quarkus

import "testing"

func TestNodeTextAndFindChild(t *testing.T) {
	root, _ := parseJava([]byte(`class C {}`))
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no class")
	}
	if findChildByType(classes[0], "class_body") == nil {
		t.Fatal("expected class_body")
	}
	if findChildByType(classes[0], "nonexistent") != nil {
		t.Fatal("expected nil")
	}
	id := findChildByType(classes[0], "identifier")
	if id != nil && nodeText(id, []byte(`class C {}`)) != "C" {
		t.Fatalf("name %q", nodeText(id, []byte(`class C {}`)))
	}
}
