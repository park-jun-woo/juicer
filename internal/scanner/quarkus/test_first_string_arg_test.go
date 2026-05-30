//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestFirstStringArg 테스트
package quarkus

import "testing"

func TestFirstStringArg(t *testing.T) {
	root, src := parseQ(t, `class C { @Path("/abc") void m() {} }`)
	ann := findAllByType(root, "annotation")
	if len(ann) == 0 {
		ann = findAllByType(root, "marker_annotation")
	}
	if got := firstStringArg(ann[0], src); got != "/abc" {
		t.Fatalf("got %q", got)
	}
}
