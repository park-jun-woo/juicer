//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what seedRouterNodes: 모든 (file,var)를 그래프 노드로 등록
package express

import "testing"

func TestSeedRouterNodes(t *testing.T) {
	g := newTestGraph()
	allRouters := map[string]map[string]bool{
		"a.ts": {"app": true},
		"b.ts": {"users": true, "admin": true},
	}
	seedRouterNodes(g, allRouters)
	want := []routerKey{
		{file: "a.ts", varName: "app"},
		{file: "b.ts", varName: "users"},
		{file: "b.ts", varName: "admin"},
	}
	for _, k := range want {
		if !g.nodes[k] {
			t.Errorf("missing node %+v", k)
		}
	}
	if len(g.nodes) != 3 {
		t.Fatalf("expected 3 nodes, got %d", len(g.nodes))
	}
}
