//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestPropagateEdge_EmptyParents 테스트
package express

import "testing"

func TestPropagateEdge_EmptyParents(t *testing.T) {
	child := routerKey{file: "b", varName: "users"}
	prefixes := map[routerKey][]string{}
	e := routerEdge{child: child, seg: "/users"}
	if propagateEdge(prefixes, e, nil) {
		t.Fatal("expected changed=false for empty parents")
	}
}
