//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestIsTypeBoxObjectCall_NotObject 테스트
package fastify

import "testing"

func TestIsTypeBoxObjectCall_NotObject(t *testing.T) {
	c, src := firstCall(t, "const x = Type.String();\n")
	if isTypeBoxObjectCall(c, src) {
		t.Error("Type.String() should not match")
	}
}
