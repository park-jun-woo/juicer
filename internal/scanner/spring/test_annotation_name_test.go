//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestAnnotationName 테스트
package spring

import "testing"

func TestAnnotationName(t *testing.T) {
	root, src := parseS(t, `class C { @org.springframework.web.bind.annotation.GetMapping void m() {} }`)
	anns := findAllByType(root, "marker_annotation")
	if len(anns) == 0 {
		t.Skip("no marker annotation")
	}
	if got := annotationName(anns[0], src); got != "GetMapping" {
		t.Fatalf("got %q", got)
	}
}
