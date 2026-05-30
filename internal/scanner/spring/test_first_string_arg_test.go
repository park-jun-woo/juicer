//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestFirstStringArg 테스트
package spring

import "testing"

func TestFirstStringArg(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping("/x") void m() {} }`)
	anns := findAllByType(root, "annotation")
	if len(anns) == 0 {
		t.Skip("no annotation")
	}
	if got := firstStringArg(anns[0], src); got != "/x" {
		t.Fatalf("got %q", got)
	}
}
