//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestAnnotationElementValue 테스트
package spring

import "testing"

func TestAnnotationElementValue(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping(value = "/x") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	if got := annotationElementValue(ann, src, "value"); got != "/x" {
		t.Fatalf("got %q", got)
	}
}
