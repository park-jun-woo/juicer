//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestAnnotationIntValue 테스트
package spring

import "testing"

func TestAnnotationIntValue(t *testing.T) {
	root, src := parseS(t, `class C { @Size(min = 3) String s; }`)
	ann := findAllByType(root, "annotation")[0]
	if v, ok := annotationIntValue(ann, src, "min"); !ok || v != 3 {
		t.Fatalf("got %d %v", v, ok)
	}
}
