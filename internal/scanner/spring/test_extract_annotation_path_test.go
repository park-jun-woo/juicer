//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractAnnotationPath 테스트
package spring

import "testing"

func TestExtractAnnotationPath(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping(path = "/p") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	if got := extractAnnotationPath(ann, src); got != "/p" {
		t.Fatalf("got %q", got)
	}
}
