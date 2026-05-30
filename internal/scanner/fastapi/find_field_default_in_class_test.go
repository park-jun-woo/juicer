//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findFieldDefaultInClass: 문자열 기본값 / 비문자열 / 필드없음 / block없음
package fastapi

import "testing"

func TestFindFieldDefaultInClass_String(t *testing.T) {
	cls, src := firstClass(t, []byte("class C:\n    name: str = \"hello\"\n"))
	if got := findFieldDefaultInClass(cls, "name", src); got != "hello" {
		t.Fatalf("got %q", got)
	}
}

func TestFindFieldDefaultInClass_NonString(t *testing.T) {
	cls, src := firstClass(t, []byte("class C:\n    count: int = 5\n"))
	if got := findFieldDefaultInClass(cls, "count", src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestFindFieldDefaultInClass_FieldMissing(t *testing.T) {
	cls, src := firstClass(t, []byte("class C:\n    x: str = \"y\"\n"))
	if got := findFieldDefaultInClass(cls, "other", src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestFindFieldDefaultInClass_NoBlock(t *testing.T) {
	// pass a node that has no block child (the module root) -> returns ""
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := findFieldDefaultInClass(root, "x", src); got != "" {
		t.Fatalf("got %q", got)
	}
}
