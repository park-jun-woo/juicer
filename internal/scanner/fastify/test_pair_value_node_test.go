//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestPairValueNode 테스트
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
