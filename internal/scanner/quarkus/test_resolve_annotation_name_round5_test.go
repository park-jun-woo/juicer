//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveAnnotationName_Round5 테스트
package quarkus

import "testing"

func TestResolveAnnotationName_Round5(t *testing.T) {
	root, src := qParse(t, `@PathParam("id") class C {}`)
	ann := qFirst(t, root, "annotation")
	if got := resolveAnnotationName(ann, src, "fallback"); got != "id" {
		t.Fatalf("got %q", got)
	}

	if got := resolveAnnotationName(nil, src, "fallback"); got != "fallback" {
		t.Fatalf("nil fallback: %q", got)
	}
}
