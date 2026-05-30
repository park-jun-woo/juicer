//ff:func feature=scan type=test control=sequence topic=django
//ff:what firstIdentArg — argument_list 첫 identifier 인자 추출 분기를 검증
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

func TestFirstIdentArg_None(t *testing.T) {
	src := []byte("x = path('a/', 'b/')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if got := firstIdentArg(args, src); got != "" {
		t.Fatalf("expected empty (only string args), got %q", got)
	}
}
