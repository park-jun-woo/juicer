//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveAnnotationName_Round5 테스트
package spring

import "testing"

func TestResolveAnnotationName_Round5(t *testing.T) {
	root, src := sParse(t, `@PathVariable("id") class C {}`)
	ann := sFirst(t, root, "annotation")
	if got := resolveAnnotationName(ann, src, "fb"); got != "id" {
		t.Fatalf("got %q", got)
	}
	if got := resolveAnnotationName(nil, src, "fb"); got != "fb" {
		t.Fatalf("nil: %q", got)
	}
}
