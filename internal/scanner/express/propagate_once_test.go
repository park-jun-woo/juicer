//ff:func feature=scan type=test control=sequence topic=express
//ff:what propagateOnce: 변경 발생 / 변경 없음
package express

import "testing"

func TestPropagateOnce_Changed(t *testing.T) {
	g := newTestGraph()
	parent := routerKey{file: "a", varName: "app"}
	child := routerKey{file: "b", varName: "users"}
	graphAddEdge(g, parent, child, "/users")
	prefixes := map[routerKey][]string{parent: {""}}

	if !propagateOnce(g, []routerKey{parent}, prefixes) {
		t.Fatal("expected changed=true")
	}
	if len(prefixes[child]) != 1 || prefixes[child][0] != "/users" {
		t.Fatalf("got %v", prefixes[child])
	}

	// second pass: no new prefixes -> no change
	if propagateOnce(g, []routerKey{parent}, prefixes) {
		t.Fatal("expected changed=false on stable pass")
	}
}
