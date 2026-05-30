//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what childrenOfType 테스트
package fastify

import "testing"

func TestChildrenOfType(t *testing.T) {
	obj, _ := firstObject(t, `{ a: 1, b: 2, c: 3 }`)
	pairs := childrenOfType(obj, "pair")
	if len(pairs) != 3 {
		t.Fatalf("expected 3 pairs, got %d", len(pairs))
	}
	// type that does not exist as a direct child -> empty
	if got := childrenOfType(obj, "function_declaration"); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}
