//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAnnotationName_Round5 테스트
package quarkus

import "testing"

func TestAnnotationName_Round5(t *testing.T) {
	root, src := qParse(t, "@GET class C {}")
	ann := qFirst(t, root, "marker_annotation")
	if got := annotationName(ann, src); got != "GET" {
		t.Fatalf("got %q", got)
	}
}
