//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindFieldDefaultInClass_NoBlock 테스트
package fastapi

import "testing"

func TestFindFieldDefaultInClass_NoBlock(t *testing.T) {

	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := findFieldDefaultInClass(root, "x", src); got != "" {
		t.Fatalf("got %q", got)
	}
}
