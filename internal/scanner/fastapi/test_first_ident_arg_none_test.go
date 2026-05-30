//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFirstIdentArg_None 테스트
package fastapi

import "testing"

func TestFirstIdentArg_None(t *testing.T) {
	src := []byte("f('literal')\n")
	root, _ := parsePython(src)
	args := findAllByType(root, "argument_list")
	if got := firstIdentArg(args[0], src); got != "" {
		t.Fatalf("got %q", got)
	}
}
