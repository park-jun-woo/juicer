//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 그래프에 부모→자식 마운트 엣지를 추가하고 노드/진입차수를 갱신한다
package express

func graphAddEdge(g *mountGraph, parent, child routerKey, seg string) {
	g.nodes[parent] = true
	g.nodes[child] = true
	g.edges[parent] = append(g.edges[parent], routerEdge{child, seg})
	g.indeg[child]++
}
