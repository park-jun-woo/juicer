//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAnnotationElementValue 테스트
package quarkus

import "testing"

func TestAnnotationElementValue(t *testing.T) {
	root, src := parseQ(t, `class C { @Path(value = "/x") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	if got := annotationElementValue(ann, src, "value"); got != "/x" {
		t.Fatalf("got %q", got)
	}
}
