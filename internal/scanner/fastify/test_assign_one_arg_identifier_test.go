//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAssignOneArg_Identifier 테스트
package fastify

import "testing"

func TestAssignOneArg_Identifier(t *testing.T) {
	n, src := firstNodeOfType(t, "f(handler);\n", "identifier")
	ri := &routeInfo{}
	assignOneArg(ri, n, src)
	if ri.Handler == "" {
		t.Fatal("identifier arg should set Handler")
	}
}
