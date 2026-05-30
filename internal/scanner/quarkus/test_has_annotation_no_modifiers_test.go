//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestHasAnnotation_NoModifiers 테스트
package quarkus

import "testing"

func TestHasAnnotation_NoModifiers(t *testing.T) {

	root, _ := parseJava([]byte(`class C { void m() {} }`))
	classes := findAllByType(root, "class_body")
	if hasAnnotation(classes[0], []byte(`class C { void m() {} }`), "NotNull") {
		t.Fatal("expected false")
	}
}
