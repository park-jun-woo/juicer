//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAssignOneArg_Default 테스트
package fastify

import "testing"

func TestAssignOneArg_Default(t *testing.T) {

	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	ri := &routeInfo{Handler: "orig"}
	assignOneArg(ri, n, src)
	if ri.Handler != "orig" || ri.Schema != nil {
		t.Fatalf("default case should leave ri unchanged, got %v", ri)
	}
}
