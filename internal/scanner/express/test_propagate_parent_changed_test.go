//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestPropagateParent_Changed 테스트
package express

import "testing"

func TestPropagateParent_Changed(t *testing.T) {
	g := newTestGraph()
	parent := routerKey{file: "a", varName: "app"}
	child := routerKey{file: "b", varName: "users"}
	graphAddEdge(g, parent, child, "/users")
	prefixes := map[routerKey][]string{parent: {"/api"}}
	if !propagateParent(g, parent, prefixes) {
		t.Fatal("expected changed=true")
	}
	if prefixes[child][0] != "/api/users" {
		t.Fatalf("got %v", prefixes[child])
	}
}
