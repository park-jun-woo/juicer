//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestIsTypeBoxObjectCall_NotType 테스트
package fastify

import "testing"

func TestIsTypeBoxObjectCall_NotType(t *testing.T) {
	c, src := firstCall(t, "const x = Other.Object({});\n")
	if isTypeBoxObjectCall(c, src) {
		t.Error("Other.Object() should not match")
	}
}
