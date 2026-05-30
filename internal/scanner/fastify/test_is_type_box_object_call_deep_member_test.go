//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestIsTypeBoxObjectCall_DeepMember 테스트
package fastify

import "testing"

func TestIsTypeBoxObjectCall_DeepMember(t *testing.T) {

	c, src := firstCall(t, "const x = Type.Object({ a: Type.Number() });\n")
	if !isTypeBoxObjectCall(c, src) {
		t.Error("expected Type.Object with nested call to match")
	}
}
