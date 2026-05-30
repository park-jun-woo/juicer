//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIsEnumClass 테스트
package spring

import "testing"

func TestIsEnumClass(t *testing.T) {
	root, _ := parseS(t, `enum E { A }`)
	en := findAllByType(root, "enum_declaration")[0]
	if !isEnumClass(en) {
		t.Fatal("enum")
	}
	root2, _ := parseS(t, `class C {}`)
	cls := findAllByType(root2, "class_declaration")[0]
	if isEnumClass(cls) {
		t.Fatal("class is not enum")
	}
}
