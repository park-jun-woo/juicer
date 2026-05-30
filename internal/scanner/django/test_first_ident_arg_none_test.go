//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestFirstIdentArg_None 테스트
package django

import "testing"

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
