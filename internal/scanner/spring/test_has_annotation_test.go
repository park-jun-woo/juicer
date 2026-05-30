//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestHasAnnotation 테스트
package spring

import "testing"

func TestHasAnnotation(t *testing.T) {
	root, src := parseS(t, `class C { @NotNull String x; }`)
	field := findAllByType(root, "field_declaration")[0]
	if !hasAnnotation(field, src, "NotNull") || hasAnnotation(field, src, "Email") {
		t.Fatal("hasAnnotation")
	}
}
