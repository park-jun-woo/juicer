//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFirstIdentArg 테스트
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
