//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestChildrenOfType 테스트
package spring

import "testing"

func TestChildrenOfType(t *testing.T) {
	root, _ := parseS(t, `class C { int a; int b; }`)
	body := findAllByType(root, "class_body")[0]
	if len(childrenOfType(body, "field_declaration")) != 2 {
		t.Fatal("children")
	}
}
