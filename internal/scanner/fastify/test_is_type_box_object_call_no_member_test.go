//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestIsTypeBoxObjectCall_NoMember 테스트
package fastify

import "testing"

func TestIsTypeBoxObjectCall_NoMember(t *testing.T) {
	c, src := firstCall(t, "const x = foo();\n")
	if isTypeBoxObjectCall(c, src) {
		t.Error("plain call should not match")
	}
}
