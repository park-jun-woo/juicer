//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFirstStringArg_None 테스트
package fastapi

import "testing"

func TestFirstStringArg_None(t *testing.T) {
	src := []byte("f(123, var)\n")
	root, _ := parsePython(src)
	args := findAllByType(root, "argument_list")
	if got := firstStringArg(args[0], src); got != "" {
		t.Fatalf("got %q", got)
	}
}
