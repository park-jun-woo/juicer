//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 모든 (file, var) 라우터 인스턴스를 그래프 노드로 등록한다
package express

func seedRouterNodes(g *mountGraph, allRouters map[string]map[string]bool) {
	for file, rs := range allRouters {
		for v := range rs {
			g.nodes[routerKey{file, v}] = true
		}
	}
}
