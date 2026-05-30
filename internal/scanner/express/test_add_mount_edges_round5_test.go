//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestAddMountEdges_Round5 테스트
package express

import "testing"

func TestAddMountEdges_Round5(t *testing.T) {
	allRouters := map[string]map[string]bool{
		"child.ts": {"sub": true, "extra": true},
	}

	g := buildMountGraph(nil, allRouters)
	addMountEdges(g, mountEntry{prefix: "/in", varName: "sub", filePath: "", sourceFile: "a.ts", sourceRouter: "r"}, allRouters)

	g2 := buildMountGraph(nil, allRouters)
	addAmbiguousMountEdges(g2, routerKey{"a.ts", "r"}, mountEntry{prefix: "/x", filePath: "child.ts", sourceFile: "a.ts", sourceRouter: "r"}, allRouters)
	if len(g2.edges) == 0 {
		t.Fatalf("expected ambiguous edges, got %v", g2.edges)
	}
}
