//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestIsApiController 테스트
package dotnet

import "testing"

func TestIsApiController(t *testing.T) {
	root, src := parseCS(t, `[ApiController] class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if !isApiController(cls, src) {
		t.Fatal("ApiController attr")
	}
	root2, src2 := parseCS(t, `class D : ControllerBase {}`)
	cls2 := findAllByType(root2, "class_declaration")[0]
	if !isApiController(cls2, src2) {
		t.Fatal("ControllerBase base")
	}
	root3, src3 := parseCS(t, `class E {}`)
	cls3 := findAllByType(root3, "class_declaration")[0]
	if isApiController(cls3, src3) {
		t.Fatal("plain class")
	}
}
