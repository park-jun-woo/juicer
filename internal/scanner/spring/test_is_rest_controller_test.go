//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIsRestController 테스트
package spring

import "testing"

func TestIsRestController(t *testing.T) {
	root, src := parseS(t, `@RestController class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if !isRestController(cls, src) {
		t.Fatal("RestController")
	}
	root2, src2 := parseS(t, `@Controller @ResponseBody class D {}`)
	cls2 := findAllByType(root2, "class_declaration")[0]
	if !isRestController(cls2, src2) {
		t.Fatal("Controller+ResponseBody")
	}
	root3, src3 := parseS(t, `class E {}`)
	cls3 := findAllByType(root3, "class_declaration")[0]
	if isRestController(cls3, src3) {
		t.Fatal("plain class")
	}
}
