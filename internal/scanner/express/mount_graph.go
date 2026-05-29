//ff:type feature=scan type=model topic=express
//ff:what 라우터 마운트 그래프(엣지/진입차수/노드 집합)
package express

type mountGraph struct {
	edges map[routerKey][]routerEdge
	indeg map[routerKey]int
	nodes map[routerKey]bool
}
