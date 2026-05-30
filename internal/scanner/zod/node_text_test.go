//ff:func feature=scan type=test control=sequence topic=zod
//ff:what nodeText 테스트 (round5)
package zod

import "testing"

func TestNodeText_Round5(t *testing.T) {
	root, src := parseTS(t, `const x = 1;`)
	got := nodeText(root, src)
	if got != "const x = 1;" {
		t.Fatalf("nodeText: got %q", got)
	}
}
