//ff:func feature=scan type=test control=sequence topic=express
//ff:what initRootPrefixes: indeg=0 루트에 빈 prefix / indeg>0 제외
package express

import "testing"

func TestInitRootPrefixes(t *testing.T) {
	g := newTestGraph()
	root := routerKey{file: "a", varName: "app"}
	child := routerKey{file: "b", varName: "users"}
	g.nodes[root] = true
	g.nodes[child] = true
	g.indeg[child] = 1 // has incoming mount

	prefixes := initRootPrefixes(g)
	if got, ok := prefixes[root]; !ok || len(got) != 1 || got[0] != "" {
		t.Fatalf("root prefix wrong: %v", prefixes)
	}
	if _, ok := prefixes[child]; ok {
		t.Fatalf("child with indeg>0 should be excluded: %v", prefixes)
	}
}
