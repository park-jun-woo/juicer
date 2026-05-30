//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAssignOneArg_Arrow 테스트
package fastify

import "testing"

func TestAssignOneArg_Arrow(t *testing.T) {
	n, src := firstNodeOfType(t, "const x = () => 1;\n", "arrow_function")
	ri := &routeInfo{}
	assignOneArg(ri, n, src)
	if ri.Handler != "(anonymous)" {
		t.Fatalf("arrow should set anonymous handler, got %q", ri.Handler)
	}
}
