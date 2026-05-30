//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAnnotationArgs_Round5 테스트
package quarkus

import "testing"

func TestAnnotationArgs_Round5(t *testing.T) {
	root, src := qParse(t, `@Path("/x") class C {}`)
	ann := qFirst(t, root, "annotation")
	args := annotationArgs(ann, src)
	if args == nil {
		t.Fatal("expected args node")
	}
}
