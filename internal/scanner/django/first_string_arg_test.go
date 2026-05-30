//ff:func feature=scan type=test control=sequence topic=django
//ff:what firstStringArg — argument_list 첫 문자열 인자 추출 분기를 검증
package django

import "testing"

func TestFirstStringArg(t *testing.T) {
	src := []byte("x = path('users/', view)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if args == nil {
		t.Fatal("no argument_list")
	}
	if got := firstStringArg(args, src); got != "users/" {
		t.Fatalf("got %q, want users/", got)
	}
}

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
