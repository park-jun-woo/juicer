//ff:func feature=scan type=test control=sequence topic=express
//ff:what newTestGraph 테스트 헬퍼
package express

func newTestGraph() *mountGraph {
	return &mountGraph{
		edges: map[routerKey][]routerEdge{},
		indeg: map[routerKey]int{},
		nodes: map[routerKey]bool{},
	}
}
