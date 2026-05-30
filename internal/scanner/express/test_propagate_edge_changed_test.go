//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestPropagateEdge_Changed 테스트
package express

import "testing"

func TestPropagateEdge_Changed(t *testing.T) {
	child := routerKey{file: "b", varName: "users"}
	prefixes := map[routerKey][]string{}
	e := routerEdge{child: child, seg: "/users"}
	if !propagateEdge(prefixes, e, []string{"/api"}) {
		t.Fatal("expected changed=true")
	}
	if len(prefixes[child]) != 1 || prefixes[child][0] != "/api/users" {
		t.Fatalf("got %v", prefixes[child])
	}
}
