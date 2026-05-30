//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what firstStringArg 테스트
package fastapi

import "testing"

func TestFirstStringArg(t *testing.T) {
	src := []byte("f('/users', 123)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "argument_list")
	if len(calls) == 0 {
		t.Fatal("no argument_list")
	}
	got := firstStringArg(calls[0], src)
	if got != "/users" {
		t.Fatalf("expected '/users', got %q", got)
	}
}

func TestFirstStringArg_None(t *testing.T) {
	src := []byte("f(123, var)\n")
	root, _ := parsePython(src)
	args := findAllByType(root, "argument_list")
	if got := firstStringArg(args[0], src); got != "" {
		t.Fatalf("got %q", got)
	}
}
