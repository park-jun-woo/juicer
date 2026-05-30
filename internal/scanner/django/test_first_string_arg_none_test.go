//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestFirstStringArg_None 테스트
package django

import "testing"

func TestFirstStringArg_None(t *testing.T) {
	src := []byte("x = path(view, other)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if got := firstStringArg(args, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
