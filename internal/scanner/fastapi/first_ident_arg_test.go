//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what firstIdentArg 테스트
package fastapi

import "testing"

func TestFirstIdentArg(t *testing.T) {
	src := []byte("f(router)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "argument_list")
	if len(calls) == 0 {
		t.Fatal("no argument_list")
	}
	got := firstIdentArg(calls[0], src)
	if got != "router" {
		t.Fatalf("expected 'router', got %q", got)
	}
}

func TestFirstIdentArg_Attribute(t *testing.T) {
	src := []byte("f(items.router)\n")
	root, _ := parsePython(src)
	args := findAllByType(root, "argument_list")
	if got := firstIdentArg(args[0], src); got != "items.router" {
		t.Fatalf("got %q", got)
	}
}

func TestFirstIdentArg_None(t *testing.T) {
	src := []byte("f('literal')\n")
	root, _ := parsePython(src)
	args := findAllByType(root, "argument_list")
	if got := firstIdentArg(args[0], src); got != "" {
		t.Fatalf("got %q", got)
	}
}
