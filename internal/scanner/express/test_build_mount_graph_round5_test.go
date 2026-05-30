//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestBuildMountGraph_Round5 테스트
package express

import "testing"

func TestBuildMountGraph_Round5(t *testing.T) {
	allRouters := map[string]map[string]bool{
		"a.ts": {"r": true},
		"b.ts": {"sub": true},
	}
	mounts := []mountEntry{
		{prefix: "/x", varName: "sub", filePath: "b.ts", sourceFile: "a.ts", sourceRouter: "r"},
	}
	g := buildMountGraph(mounts, allRouters)
	if g == nil {
		t.Fatal("nil graph")
	}
	if !g.nodes[routerKey{file: "a.ts", varName: "r"}] {
		t.Fatalf("expected seeded node, got %v", g.nodes)
	}
}
