//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFirstStringArg_NotFound 테스트
package nestjs

import "testing"

func TestFirstStringArg_NotFound(t *testing.T) {
	src := []byte(`f(42)`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	args := findAllByType(root, "arguments")
	if len(args) == 0 {
		t.Fatal("expected arguments node")
	}
	_, ok := firstStringArg(args[0], src)
	if ok {
		t.Fatal("expected false")
	}
}
