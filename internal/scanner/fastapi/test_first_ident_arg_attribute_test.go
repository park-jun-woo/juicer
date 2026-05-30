//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFirstIdentArg_Attribute 테스트
package fastapi

import "testing"

func TestFirstIdentArg_Attribute(t *testing.T) {
	src := []byte("f(items.router)\n")
	root, _ := parsePython(src)
	args := findAllByType(root, "argument_list")
	if got := firstIdentArg(args[0], src); got != "items.router" {
		t.Fatalf("got %q", got)
	}
}
