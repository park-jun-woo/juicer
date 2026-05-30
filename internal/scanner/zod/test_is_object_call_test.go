//ff:func feature=scan type=test control=iteration dimension=1 topic=zod
//ff:what TestIsObjectCall 테스트
package zod

import "testing"

func TestIsObjectCall(t *testing.T) {
	root, src := parseTS(t, `z.object({ a: z.string() });`)
	calls := findAllByType(root, "call_expression")
	found := false
	for _, c := range calls {
		if IsObjectCall(c, src) {
			found = true
		}
	}
	if !found {
		t.Fatal("expected z.object call")
	}
	root2, src2 := parseTS(t, `other.foo();`)
	calls2 := findAllByType(root2, "call_expression")
	if IsObjectCall(calls2[0], src2) {
		t.Fatal("non-z call")
	}
}
