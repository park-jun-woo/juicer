//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindFieldDefaultInClass_NonString 테스트
package fastapi

import "testing"

func TestFindFieldDefaultInClass_NonString(t *testing.T) {
	cls, src := firstClass(t, []byte("class C:\n    count: int = 5\n"))
	if got := findFieldDefaultInClass(cls, "count", src); got != "" {
		t.Fatalf("got %q", got)
	}
}
