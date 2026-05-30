//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestChildrenOfType 테스트
package hono

import "testing"

func TestChildrenOfType(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: 1, b: 2, c: 3 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	pairs := childrenOfType(obj, "pair")
	if len(pairs) != 3 {
		t.Fatalf("expected 3 pairs, got %d", len(pairs))
	}
}
