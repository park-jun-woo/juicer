//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestPropagateEdge_NoChangeOnDuplicate 테스트
package express

import "testing"

func TestPropagateEdge_NoChangeOnDuplicate(t *testing.T) {
	child := routerKey{file: "b", varName: "users"}
	prefixes := map[routerKey][]string{child: {"/api/users"}}
	e := routerEdge{child: child, seg: "/users"}
	if propagateEdge(prefixes, e, []string{"/api"}) {
		t.Fatal("expected changed=false for duplicate")
	}
}
