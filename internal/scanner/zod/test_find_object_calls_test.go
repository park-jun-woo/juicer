//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestFindObjectCalls 테스트
package zod

import "testing"

func TestFindObjectCalls(t *testing.T) {
	root, src := parseTS(t, `z.object({ a: z.string() });`)
	calls := FindObjectCalls(root, src)
	if len(calls) == 0 {
		t.Fatal("expected object calls")
	}
}
