//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 라우터 집합과 마운트 엔트리로부터 마운트 그래프(엣지/진입차수/노드)를 구축한다
package express

func buildMountGraph(mounts []mountEntry, allRouters map[string]map[string]bool) *mountGraph {
	g := &mountGraph{
		edges: map[routerKey][]routerEdge{},
		indeg: map[routerKey]int{},
		nodes: map[routerKey]bool{},
	}
	seedRouterNodes(g, allRouters)
	for _, m := range mounts {
		addMountEdges(g, m, allRouters)
	}
	return g
}
