//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestFindAnnotation 테스트
package spring

import "testing"

func TestFindAnnotation(t *testing.T) {
	field, src := firstField(t, `class C { @Size(min = 1) private String s; }`)
	if findAnnotation(field, src, "Size") == nil || findAnnotation(field, src, "Missing") != nil {
		t.Fatal("findAnnotation")
	}
}
