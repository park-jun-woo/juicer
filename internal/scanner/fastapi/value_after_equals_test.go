//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what valueAfterEquals 테스트
package fastapi

import "testing"

func TestValueAfterEquals(t *testing.T) {
	src := []byte("f(x=42)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kws := findAllByType(root, "keyword_argument")
	if len(kws) == 0 {
		t.Fatal("no keyword_argument")
	}
	got := valueAfterEquals(kws[0], src)
	if got != "42" {
		t.Fatalf("expected '42', got %q", got)
	}
}

func TestValueAfterEquals_NoEquals(t *testing.T) {
	// a node with no "=" child -> ""
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
