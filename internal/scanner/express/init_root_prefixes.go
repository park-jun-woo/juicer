//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 들어오는 마운트가 없는 루트 라우터에 prefix ""를 부여한다
package express

func initRootPrefixes(g *mountGraph) map[routerKey][]string {
	prefixes := map[routerKey][]string{}
	for n := range g.nodes {
		if g.indeg[n] == 0 {
			prefixes[n] = []string{""}
		}
	}
	return prefixes
}
