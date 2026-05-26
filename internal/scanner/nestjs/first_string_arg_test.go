//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFirstStringArg_Found 테스트
package nestjs

import "testing"

func TestFirstStringArg_Found(t *testing.T) {
	src := []byte(`f('hello')`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	args := findAllByType(root, "arguments")
	if len(args) == 0 {
		t.Fatal("expected arguments node")
	}
	val, ok := firstStringArg(args[0], src)
	if !ok || val != "hello" {
		t.Fatalf("expected hello, got %q ok=%v", val, ok)
	}
}
