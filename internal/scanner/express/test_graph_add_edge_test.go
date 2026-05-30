//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestGraphAddEdge 테스트
package express

import "testing"

func TestGraphAddEdge(t *testing.T) {
	g := newTestGraph()
	parent := routerKey{file: "a.ts", varName: "app"}
	child := routerKey{file: "b.ts", varName: "users"}
	graphAddEdge(g, parent, child, "/api")

	if !g.nodes[parent] || !g.nodes[child] {
		t.Fatalf("nodes not registered: %v", g.nodes)
	}
	if len(g.edges[parent]) != 1 || g.edges[parent][0].child != child || g.edges[parent][0].seg != "/api" {
		t.Fatalf("edge wrong: %v", g.edges[parent])
	}
	if g.indeg[child] != 1 {
		t.Fatalf("indeg=%d", g.indeg[child])
	}

	graphAddEdge(g, parent, child, "/v2")
	if g.indeg[child] != 2 || len(g.edges[parent]) != 2 {
		t.Fatalf("second edge: indeg=%d edges=%d", g.indeg[child], len(g.edges[parent]))
	}
}
