//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAssignOneArg_Object 테스트
package fastify

import "testing"

func TestAssignOneArg_Object(t *testing.T) {
	n, src := firstNodeOfType(t, "const x = { a: 1 };\n", "object")
	ri := &routeInfo{}
	assignOneArg(ri, n, src)
	if ri.Schema != n {
		t.Fatal("object arg should set Schema")
	}
}
