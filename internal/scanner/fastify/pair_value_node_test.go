//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what pairValueNode 테스트
package fastify

import "testing"

func TestPairValueNode(t *testing.T) {
	pairs, src := schemaPairs(t, `{ a: "x" }`)
	val := pairValueNode(pairs[0])
	if val == nil || val.Type() != "string" {
		t.Fatalf("expected string value, got %v", val)
	}
	_ = src
}

func TestPairValueNode_NoColon(t *testing.T) {
	// a non-pair node (object literal) has no ":" at top level -> nil
	obj, _ := firstObject(t, `{ a: 1 }`)
	if got := pairValueNode(obj); got != nil {
		t.Fatalf("expected nil for node without colon, got %s", got.Type())
	}
}
