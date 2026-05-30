//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestHasAnnotation 테스트
package quarkus

import "testing"

func TestHasAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @NotNull private String name; }`)
	if !hasAnnotation(field, src, "NotNull") {
		t.Fatal("expected NotNull")
	}
	if hasAnnotation(field, src, "Email") {
		t.Fatal("unexpected Email")
	}
}
