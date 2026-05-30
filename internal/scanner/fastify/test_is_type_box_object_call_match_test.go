//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestIsTypeBoxObjectCall_Match 테스트
package fastify

import "testing"

func TestIsTypeBoxObjectCall_Match(t *testing.T) {
	c, src := firstCall(t, "const x = Type.Object({});\n")
	if !isTypeBoxObjectCall(c, src) {
		t.Error("expected Type.Object() to match")
	}
}
