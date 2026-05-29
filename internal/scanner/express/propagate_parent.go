//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 한 부모 라우터의 prefix들을 자식 엣지로 전파하고 변경 여부를 반환한다
package express

func propagateParent(g *mountGraph, p routerKey, prefixes map[routerKey][]string) bool {
	pps, ok := prefixes[p]
	if !ok {
		return false
	}
	changed := false
	for _, e := range g.edges[p] {
		if propagateEdge(prefixes, e, pps) {
			changed = true
		}
	}
	return changed
}
