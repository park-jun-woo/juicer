//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestContainsCall 테스트
package zod

import "testing"

func TestContainsCall(t *testing.T) {
	root, src := parseTS(t, `const s = z.string().min(1);`)
	if !ContainsCall(root, src) {
		t.Fatal("expected z call")
	}
	root2, src2 := parseTS(t, `const x = 1;`)
	if ContainsCall(root2, src2) {
		t.Fatal("no z call")
	}
}
