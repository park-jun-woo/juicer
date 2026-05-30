//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAnnotationIntValue 테스트
package quarkus

import "testing"

func TestAnnotationIntValue(t *testing.T) {
	root, src := parseQ(t, `class C { @Size(min = 1, max = 10) String s; }`)
	ann := findAllByType(root, "annotation")[0]
	if v, ok := annotationIntValue(ann, src, "min"); !ok || v != 1 {
		t.Fatalf("min: %d %v", v, ok)
	}
	if _, ok := annotationIntValue(ann, src, "missing"); ok {
		t.Fatal("missing key")
	}
}
