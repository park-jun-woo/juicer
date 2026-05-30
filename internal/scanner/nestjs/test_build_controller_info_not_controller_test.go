//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildControllerInfo_NotController 테스트
package nestjs

import "testing"

func TestBuildControllerInfo_NotController(t *testing.T) {
	src := []byte(`export class PlainClass {}`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	if _, ok := buildControllerInfo(cls, src, "x.ts", "/abs/x.ts", map[string]string{}); ok {
		t.Fatal("expected not ok for non-controller")
	}
}
