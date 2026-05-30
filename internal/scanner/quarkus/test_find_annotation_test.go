//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestFindAnnotation 테스트
package quarkus

import "testing"

func TestFindAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Size(min = 1) private String s; }`)
	if findAnnotation(field, src, "Size") == nil {
		t.Fatal("expected Size annotation")
	}
	if findAnnotation(field, src, "Missing") != nil {
		t.Fatal("expected nil")
	}
}
