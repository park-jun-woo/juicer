//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindFieldDefaultInClass_FieldMissing 테스트
package fastapi

import "testing"

func TestFindFieldDefaultInClass_FieldMissing(t *testing.T) {
	cls, src := firstClass(t, []byte("class C:\n    x: str = \"y\"\n"))
	if got := findFieldDefaultInClass(cls, "other", src); got != "" {
		t.Fatalf("got %q", got)
	}
}
