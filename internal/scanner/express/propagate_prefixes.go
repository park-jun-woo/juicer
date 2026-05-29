//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 마운트 그래프를 따라 prefix가 수렴할 때까지 결정적으로 전파한다
package express

func propagatePrefixes(g *mountGraph, prefixes map[routerKey][]string) {
	parents := sortedParents(g)
	const maxIter = 100
	for it := 0; it < maxIter; it++ {
		if !propagateOnce(g, parents, prefixes) {
			break
		}
	}
}
