//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestFirstStringArg 테스트
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
