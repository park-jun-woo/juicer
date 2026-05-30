//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestPairValueNode_NoColon 테스트
package fastify

import "testing"

func TestPairValueNode_NoColon(t *testing.T) {

	obj, _ := firstObject(t, `{ a: 1 }`)
	if got := pairValueNode(obj); got != nil {
		t.Fatalf("expected nil for node without colon, got %s", got.Type())
	}
}
