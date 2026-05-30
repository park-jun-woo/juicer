//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestMatchAnnotationRoute_GetMapping 테스트
package spring

import "testing"

func TestMatchAnnotationRoute_GetMapping(t *testing.T) {
	root, src := parseS(t, `class C { @GetMapping("/x") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	method, path, ok := matchAnnotationRoute(ann, src)
	if !ok || method != "GET" || path != "/x" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}
