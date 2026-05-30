//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestMatchAnnotationRoute_RequestMapping 테스트
package spring

import "testing"

func TestMatchAnnotationRoute_RequestMapping(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping(value = "/x", method = RequestMethod.POST) void m() {} }`)
	anns := findAllByType(root, "annotation")
	method, path, ok := matchAnnotationRoute(anns[0], src)
	if !ok || path != "/x" || method != "POST" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}
