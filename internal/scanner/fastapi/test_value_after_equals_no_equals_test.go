//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestValueAfterEquals_NoEquals 테스트
package fastapi

import "testing"

func TestValueAfterEquals_NoEquals(t *testing.T) {

	src := []byte("f(positional)\n")
	root, _ := parsePython(src)
	args := findAllByType(root, "argument_list")
	if len(args) == 0 {
		t.Fatal("no argument_list")
	}
	if got := valueAfterEquals(args[0], src); got != "" {
		t.Fatalf("expected empty for node without '=', got %q", got)
	}
}
