//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestFirstIdentArg 테스트
package django

import "testing"

func TestFirstIdentArg(t *testing.T) {
	src := []byte("x = path('a/', my_view)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if args == nil {
		t.Fatal("no argument_list")
	}
	if got := firstIdentArg(args, src); got != "my_view" {
		t.Fatalf("got %q, want my_view", got)
	}
}
