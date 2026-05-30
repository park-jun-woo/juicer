//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestSingleIntArg 테스트
package spring

import "testing"

func TestSingleIntArg(t *testing.T) {
	root, src := parseS(t, `class C { @Min(7) int x; }`)
	ann := findAllByType(root, "annotation")[0]
	if v, ok := singleIntArg(ann, src); !ok || v != 7 {
		t.Fatalf("got %d %v", v, ok)
	}
}
