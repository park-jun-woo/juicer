//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindFieldDefaultInClass_String 테스트
package fastapi

import "testing"

func TestFindFieldDefaultInClass_String(t *testing.T) {
	cls, src := firstClass(t, []byte("class C:\n    name: str = \"hello\"\n"))
	if got := findFieldDefaultInClass(cls, "name", src); got != "hello" {
		t.Fatalf("got %q", got)
	}
}
