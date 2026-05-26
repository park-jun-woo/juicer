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
