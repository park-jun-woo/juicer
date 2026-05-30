//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestAnnotationArgs_Round5 테스트
package spring

import "testing"

func TestAnnotationArgs_Round5(t *testing.T) {
	root, src := sParse(t, `@RequestMapping("/x") class C {}`)
	ann := sFirst(t, root, "annotation")
	if annotationArgs(ann, src) == nil {
		t.Fatal("expected args")
	}
}
